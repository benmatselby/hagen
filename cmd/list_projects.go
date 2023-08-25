package cmd

import (
	"fmt"
	"io"
	"os"

	hagen "github.com/benmatselby/hagen/pkg"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ListProjectsOptions defines what arguments/options the user can provide for the `project ls` command
type ListProjectsOptions struct {
	Args  []string
	Org   string
	Repo  string
	State string
}

// NewListProjectsCommand creates a new `projects` command
func NewListProjectsCommand(client hagen.Provider) *cobra.Command {
	var opts ListProjectsOptions

	cmd := &cobra.Command{
		Use:   "projects",
		Short: "List the projects. Default query is to list project for ${GITHUB_OWNER}",
		RunE: func(cmd *cobra.Command, args []string) error {
			return ListProjects(client, opts, os.Stdout)
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&opts.Org, "org", "", "The organisation to get projects for")
	flags.StringVar(&opts.Repo, "repo", "", "The repo to get projects for")
	flags.StringVar(&opts.State, "state", "", "What state of projects to show: open (default), closed, or all")

	return cmd
}

// ListProjects will display projects based on the given search criteria
func ListProjects(client hagen.Provider, userOpts ListProjectsOptions, w io.Writer) error {
	states := map[string]bool{
		"open":   true,
		"closed": true,
		"all":    true,
	}

	if !states[userOpts.State] {
		userOpts.State = "open"
	}

	if userOpts.Org == "" && userOpts.Repo == "" {
		userOpts.Org = viper.GetString("GITHUB_ORG")
	}

	opts := github.ProjectListOptions{
		State: userOpts.State,
	}

	if userOpts.Org != "" {
		projects, _, err := client.ListProjectsForOrg(userOpts.Org, opts)
		if err != nil {
			return err
		}
		DisplayProjects(projects, w)
	} else if userOpts.Repo != "" {
		projects, _, err := client.ListProjectsForRepo(userOpts.Repo, opts)
		if err != nil {
			return err
		}
		DisplayProjects(projects, w)
	} else {
		projects, _, err := client.ListProjectsForOrg(viper.GetString("GITHUB_OWNER"), opts)
		if err != nil {
			return err
		}
		DisplayProjects(projects, w)
	}

	return nil
}

// DisplayProjects will render the projects
func DisplayProjects(projects []*github.Project, w io.Writer) {
	for _, project := range projects {
		fmt.Fprintln(w, project.GetName())
	}
}
