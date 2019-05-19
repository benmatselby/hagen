package pkg

import (
	"github.com/google/go-github/github"
)

// Issue represents an issue in we want to display in the app
type Issue struct {
	Ref    string
	Title  string
	Points int
}

// Provider is the interface to the back end data source
type Provider interface {
	ListIssues(query string, opts github.SearchOptions) (*github.IssuesSearchResult, error)
}
