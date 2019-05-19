package pkg

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

// Client is the GitHub wrapper and concrete implementation or Provider
type Client struct {
	GH      *github.Client
	Context context.Context
	User    string
	Token   string
}

// New will provide a GitHub client
func New() Client {
	token := viper.GetString("GITHUB_TOKEN")
	owner := viper.GetString("GITHUB_OWNER")

	ctx := context.Background()
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	http := oauth2.NewClient(ctx, tokenSource)
	ghClient := github.NewClient(http)

	client := Client{
		GH:      ghClient,
		Context: ctx,
		User:    owner,
		Token:   token,
	}

	return client
}

// ListIssues will return issues from GitHub
func (client *Client) ListIssues(query string, opts github.SearchOptions) (*github.IssuesSearchResult, error) {
	result, _, err := client.GH.Search.Issues(client.Context, query, &opts)
	if err != nil {
		return nil, err
	}

	return result, err
}
