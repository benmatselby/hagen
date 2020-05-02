package pkg

import (
	"context"
	"strings"
	"sync"

	"github.com/google/go-github/github"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

// Client is the GitHub wrapper and concrete implementation or Provider
type Client struct {
	GH      *github.Client
	Context context.Context
	User    string
	Org     string
	Token   string
}

// New will provide a GitHub client
func New() Client {
	token := viper.GetString("GITHUB_TOKEN")
	owner := viper.GetString("GITHUB_OWNER")
	org := viper.GetString("GITHUB_ORG")

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
		Org:     org,
		Token:   token,
	}

	return client
}

// ListIssues will return issues from GitHub
func (c *Client) ListIssues(query string, opts github.SearchOptions) (*github.IssuesSearchResult, error) {
	result, _, err := c.GH.Search.Issues(c.Context, query, &opts)
	if err != nil {
		return nil, err
	}

	return result, err
}

// ListRepos will return repos from GitHub for the given organisation in context
// See https://help.github.com/en/github/searching-for-information-on-github/searching-for-repositories for more information.
func (c *Client) ListRepos(query string, opts github.SearchOptions) (*github.RepositoriesSearchResult, error) {
	result, _, err := c.GH.Search.Repositories(c.Context, query, &opts)
	if err != nil {
		return nil, err
	}

	return result, err
}

// ListProjectsForOrg will return the projects defined against an organisation
func (c *Client) ListProjectsForOrg(orgName string, opts github.ProjectListOptions) ([]*github.Project, *github.Response, error) {
	result, res, err := c.GH.Organizations.ListProjects(c.Context, orgName, &opts)
	if err != nil {
		return nil, nil, err
	}

	return result, res, err
}

// ListProjectsForRepo will return the projects defined against a repo
func (c *Client) ListProjectsForRepo(repoName string, opts github.ProjectListOptions) ([]*github.Project, *github.Response, error) {
	parts := strings.Split(repoName, "/")

	var user, repo string
	if len(parts) < 2 {
		user = c.User
		repo = parts[0]
	} else {
		user = parts[0]
		repo = parts[1]
	}

	result, res, err := c.GH.Repositories.ListProjects(c.Context, user, repo, &opts)

	if err != nil {
		return nil, nil, err
	}

	return result, res, err
}

// ListColumnsForProject will return columns for a project board
func (c *Client) ListColumnsForProject(projectName, org, repo string) ([]*github.ProjectColumn, error) {
	project := c.GetProjectByName(projectName, org, repo)
	opts := github.ListOptions{}
	columns, _, err := c.GH.Projects.ListProjectColumns(c.Context, project.GetID(), &opts)
	if err != nil {
		return nil, err
	}

	return columns, nil
}

// ListIssuesForProjectColumn will return a slice of issues for a given project column
func (c *Client) ListIssuesForProjectColumn(columnID int64) ([]*github.Issue, error) {
	opts := github.ProjectCardListOptions{}
	var cards []*github.ProjectCard

	for {
		projectCards, response, err := c.GH.Projects.ListProjectCards(c.Context, columnID, &opts)
		if err != nil {
			return nil, err
		}

		cards = append(cards, projectCards...)
		if response.NextPage == 0 {
			break
		}
		opts.Page = response.NextPage
	}

	ch := make(chan *github.Issue)
	var wg sync.WaitGroup
	wg.Add(len(cards))

	go func() {
		wg.Wait()
		close(ch)
	}()

	for _, card := range cards {
		go func(card *github.ProjectCard) {
			defer wg.Done()
			request, err := c.GH.NewRequest("GET", card.GetContentURL(), nil)
			if err != nil {
				return
			}

			var issue *github.Issue
			_, err = c.GH.Do(c.Context, request, &issue)
			if err != nil {
				return
			}
			ch <- issue
		}(card)
	}

	var issues []*github.Issue
	for issue := range ch {
		issues = append(issues, issue)
	}

	return issues, nil
}

// GetProjectByName will return a single project given a name
func (c *Client) GetProjectByName(name, org, repo string) *github.Project {
	var projects []*github.Project

	if org != "" {
		opts := github.ProjectListOptions{}
		projects, _, _ = c.ListProjectsForOrg(org, opts)
	} else if repo != "" {
		opts := github.ProjectListOptions{}
		projects, _, _ = c.ListProjectsForRepo(repo, opts)
	}

	for _, project := range projects {
		if project.GetName() == name {
			return project
		}
	}

	return nil
}
