package k8s

import (
	"github.com/jumale/alfred-go/pkg/alfred"
	"strings"
	"text/template"
)

type ContextsFilter struct {
	Client
	Config
}

func (f *ContextsFilter) Filter(query string, arg *template.Template) (*alfred.List, error) {
	var err error
	curr := Item{}

	current, err := f.GetCurrentContext()
	if err != nil {
		return nil, err
	}

	var items []alfred.ListItem
	contexts, err := f.Contexts()
	if err != nil {
		return nil, err
	}

	for _, curr.Context = range contexts {
		ctxTitle := f.ContextTitle(curr.Context)

		// filter by query, if specified
		if query != "" && !strings.Contains(curr.Context, query) {
			continue
		}

		prefix := ""
		if curr.Context == current {
			prefix = "* "
		}

		items = append(items, alfred.ListItem{
			Title:    prefix + ctxTitle,
			Uid:      curr.Context,
			Arg:      generateArg(arg, curr),
			Subtitle: curr.Context,
			Order:    f.ContextOrder(curr.Context),
		})
	}

	return &alfred.List{Items: items}, nil
}
