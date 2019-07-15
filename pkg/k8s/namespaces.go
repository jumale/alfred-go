package k8s

import (
	"github.com/jumale/alfred-go/pkg/alfred"
	"strings"
	"text/template"
)

type NamespacesFilter struct {
	Client
	Config
}

func (f *NamespacesFilter) Filter(query string, arg *template.Template) (*alfred.List, error) {
	var err error
	curr := Item{}

	current, err := f.GetCurrentNamespace()
	if err != nil {
		return nil, err
	}

	var items []alfred.ListItem
	namespaces, err := f.Namespaces()
	if err != nil {
		return nil, err
	}

	for _, curr.Namespace = range namespaces {
		nsTitle := f.ContextTitle(curr.Namespace)

		// filter by query, if specified
		if query != "" && !strings.Contains(curr.Namespace, query) {
			continue
		}

		prefix := ""
		if curr.Namespace == current {
			prefix = "* "
		}

		items = append(items, alfred.ListItem{
			Title:    prefix + nsTitle,
			Uid:      curr.Namespace,
			Arg:      generateArg(arg, curr),
			Subtitle: curr.Namespace,
		})
	}

	return &alfred.List{Items: items}, nil
}
