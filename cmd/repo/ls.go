package repo

import (
	"fmt"
	"io"
	"os"

	hagen "github.com/benmatselby/hagen/pkg"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// LsOptions defines what arguments/options the user can provide for the `repo ls` command
type LsOptions struct {
	Args     []string
	Count    int
	Query    string
	Template string
	Verbose  bool
}

// NewRepoLsCommand creates a new `repo ls` command that lists the repos for the
// organisation
func NewRepoLsCommand(client hagen.Provider) *cobra.Command {
	var opts LsOptions

	cmd := &cobra.Command{
		Use:   "ls",
		Short: "List the repositories based on a query. Default query is to list repos by ${GITHUB_OWNER}",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Args = args

			query, searchOpts := NewSearchFromRepoOptions(opts)
			result, err := client.ListRepos(query, searchOpts)
			if err != nil {
				return err
			}

			err = DisplayRepos(result, os.Stdout)
			if err != nil {
				return err
			}

			if opts.Verbose {
				fmt.Fprintf(os.Stdout, "\n\nQuery used: %s\n", query)
			}

			return nil
		},
	}

	flags := cmd.Flags()
	flags.IntVar(&opts.Count, "count", -1, "How many repositories to display")
	flags.StringVar(&opts.Query, "query", "", "The search query to get repositories")
	flags.StringVar(&opts.Template, "template", "", "Use a query defined in the configuration file")
	flags.BoolVar(&opts.Verbose, "verbose", false, "Produce verbose output")

	return cmd
}

// NewSearchFromRepoOptions will return details that can perform a GitHub search
func NewSearchFromRepoOptions(opts LsOptions) (string, github.SearchOptions) {
	query := ""
	searchOpts := github.SearchOptions{}

	if opts.Template != "" && viper.IsSet(fmt.Sprintf("templates.%s", opts.Template)) {
		count := viper.GetInt(fmt.Sprintf("templates.%s.count", opts.Template))
		query = viper.GetString(fmt.Sprintf("templates.%s.query", opts.Template))
		searchOpts = github.SearchOptions{
			ListOptions: github.ListOptions{PerPage: count},
		}
	} else if opts.Query != "" {
		if opts.Query != "" {
			query = opts.Query
		}
	}

	if opts.Count != -1 {
		searchOpts.ListOptions.PerPage = opts.Count
	}

	if query == "" {
		query = fmt.Sprintf("org:%s", viper.GetString("GITHUB_OWNER"))
	}

	return query, searchOpts
}

// DisplayRepos will display issues based on the given search criteria
func DisplayRepos(result *github.RepositoriesSearchResult, w io.Writer) error {
	for _, repo := range result.Repositories {
		fmt.Println(repo.GetFullName())
	}

	return nil
}
