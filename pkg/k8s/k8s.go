package k8s

import (
	"bytes"
	"text/template"
)

type Client interface {
	GetCurrentContext() (string, error)
	SetCurrentContext(name string) error

	GetCurrentNamespace() (string, error)
	SetCurrentNamespace(name string) error

	Contexts() ([]string, error)
	Namespaces() ([]string, error)
	Pods() ([]string, error)
	Containers(pod string) ([]string, error)

	PodExec(pod string, container string, opts string, cmd string) error
	PodLogs(pod string, container string) error
	PodForwardPorts(pod string, localPort int, remotePort int) error
}

type Item struct {
	Pod       string
	Container string
	Namespace string
	Context   string
}

type Config struct {
	BinPath        string
	PodTitle       TitleFunc
	ContextTitle   TitleFunc
	ContainerTitle TitleFunc
	NamespaceTitle TitleFunc
	ContextOrder   OrderFunc
}

type TitleFunc func(name string) (title string)
type OrderFunc func(name string) (order int)

func DefaultTitle(name string) string {
	return name
}

func DefaultOrder(_ string) int {
	return 0
}

func generateArg(tpl *template.Template, vars Item) string {
	argBuf := &bytes.Buffer{}
	err := tpl.Execute(argBuf, vars)
	if err != nil {
		panic(err)
	}
	return argBuf.String()
}
