package cli

import (
	"os/exec"
	"strings"
)

type Client struct {
	BinPath string
}

func (c *Client) exec(args string) (string, error) {
	out, err := exec.Command(c.BinPath, strings.Fields(args)...).Output()
	return string(out), err
}

func (c *Client) execList(args string) ([]string, error) {
	out, err := c.exec(args)
	return strings.Split(out, `\n`), err
}

func (c *Client) GetCurrentContext() (string, error) {
	return c.exec(CmdGetCurrentContext)
}

func (c *Client) SetCurrentContext(name string) error {
	_, err := c.exec(CmdSetCurrentContext(name))
	return err
}

func (c *Client) GetCurrentNamespace() (string, error) {
	return c.exec(CmdGetCurrentNamespace)
}

func (c *Client) SetCurrentNamespace(name string) error {
	_, err := c.exec(CmdSetCurrentNamespace(c.BinPath, name))
	return err
}

func (c *Client) Contexts() ([]string, error) {
	return c.execList(CmdGetContexts)
}

func (c *Client) Namespaces() ([]string, error) {
	list, err := c.execList(CmdGetNamespaces)
	if err != nil {
		return list, err
	}
	for key, ns := range list {
		list[key] = strings.Replace(ns, "namespaces/", "", 1)
	}
	return list, nil
}

func (c *Client) Pods() ([]string, error) {
	list, err := c.execList(CmdGetPods)
	if err != nil {
		return list, err
	}
	for key, ns := range list {
		list[key] = strings.Replace(ns, "pods/", "", 1)
	}
	return list, nil
}

func (c *Client) Containers(pod string) ([]string, error) {
	out, err := c.exec(CmdPodContainers(pod))
	return strings.Split(out, " "), err
}

func (c *Client) PodExec(pod string, container string, opts string, cmd string) error {
	_, err := c.exec(CmdPodExec(pod, container, opts, cmd))
	return err
}

func (c *Client) PodLogs(pod string, container string) error {
	_, err := c.exec(CmdPodLogs(pod, container))
	return err
}

func (c *Client) PodForwardPorts(pod string, localPort int, remotePort int) error {
	_, err := c.exec(CmdPodForwardPorts(pod, localPort, remotePort))
	return err
}
