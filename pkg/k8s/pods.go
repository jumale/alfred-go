package k8s

import (
	"fmt"
	"github.com/jumale/alfred-go/pkg/alfred"
	"strings"
	"text/template"
)

type PodsConfig struct {
	Title          TitleFunc
	ContainerTitle TitleFunc
}

type PodsFilter struct {
	Client
	Config
}

func (f *PodsFilter) Filter(query string, arg *template.Template) (*alfred.List, error) {
	var err error
	curr := Item{}

	curr.Context, err = f.GetCurrentContext()
	if err != nil {
		return nil, err
	}

	curr.Namespace, err = f.GetCurrentNamespace()
	if err != nil {
		return nil, err
	}

	ctxTitle := f.ContextTitle(curr.Context)
	nsTitle := f.ContextTitle(curr.Namespace)

	var items []alfred.ListItem
	pods, err := f.Pods()
	if err != nil {
		return nil, err
	}

	for _, curr.Pod = range pods {
		// filter by query, if specified
		if query != "" && !strings.Contains(curr.Pod, query) {
			continue
		}

		podTitle := f.PodTitle(curr.Pod)
		containers, err := f.Containers(curr.Pod)
		if err != nil {
			return nil, err
		}

		for _, curr.Container = range containers {
			items = append(items, alfred.ListItem{
				Title:    podTitle,
				Uid:      curr.Pod,
				Arg:      generateArg(arg, curr),
				Subtitle: fmt.Sprintf("%s / %s / %s", f.ContainerTitle(curr.Container), nsTitle, ctxTitle),
			})
		}
	}

	return &alfred.List{Items: items}, nil
}
