package alfred

import "text/template"

type List struct {
	Items []ListItem `json:"items"`
}

type ListItem struct {
	Title    string `json:"title,omitempty"`
	Subtitle string `json:"subtitle,omitempty"`
	Uid      string `json:"uid,omitempty"`
	Arg      string `json:"arg,omitempty"`
	Order    int    `json:"order,omitempty"`
}

// Represents a service which returns an Alfred List.
type ListFilter interface {
	// Filter returns a list of found items.
	// If 'query' is not empty, then the items should be filtered where title is
	// matching the query.
	// The 'arg' provides a compiled template, which is used to generate 'Arg'
	// parameter if a ListItem. Every implementation suppose to pass to the
	// 'arg' template some specific set of variables which are specific for the
	// current implementation.
	Filter(query string, arg *template.Template) (*List, error)
}
