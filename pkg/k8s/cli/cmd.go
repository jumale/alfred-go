package cli

import (
	"fmt"
)

const (
	CmdGetCurrentContext   = "config current-context"
	CmdGetCurrentNamespace = "config view --minify=true -o jsonpath='{.contexts[0].context.namespace}'"
	CmdGetContexts         = "config get-contexts -o name"
	CmdGetNamespaces       = "get namespaces -o name"
	CmdGetPods             = "get pods -o name"
)

func CmdSetCurrentContext(ctx string) string {
	return fmt.Sprintf("config use-context %s", ctx)
}

func CmdSetCurrentNamespace(kubectlPath string, ns string) string {
	return fmt.Sprintf("config set-context $(%s config current-context) --namespace=%s", kubectlPath, ns)
}

func CmdPodExec(pod string, container string, opts string, cmd string) string {
	if container != "" {
		container = "--container=" + container
	}
	return fmt.Sprintf("exec %s %s %s %s", opts, container, pod, cmd)
}

func CmdPodLogs(pod string, container string) string {
	if container != "" {
		container = "--container=" + container
	}
	return fmt.Sprintf("logs -f %s %s", container, pod)
}

func CmdPodForwardPorts(pod string, localPort int, remotePort int) string {
	return fmt.Sprintf("port-forward %s %d:%d", pod, localPort, remotePort)
}

func CmdPodContainers(pod string) string {
	return fmt.Sprintf("get pods %s -o jsonpath='{.spec.containers[*].name}'", pod)
}
