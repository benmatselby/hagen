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
	ListProjectsForOrg(orgName string, opts github.ProjectListOptions) ([]*github.Project, *github.Response, error)
	ListProjectsForRepo(repoName string, opts github.ProjectListOptions) ([]*github.Project, *github.Response, error)
	ListRepos(query string, opts github.SearchOptions) (*github.RepositoriesSearchResult, error)
}

//go:generate mockgen -source=provider.go -package=pkg -destination=mock_github.go
