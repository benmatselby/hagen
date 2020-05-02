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

// ListRepoOptions defines what arguments/options the user can provide for
// the `repos` command.
type ListRepoOptions struct {
	Args      []string
	Count     int
	Query     string
	Recursive bool
	Template  string
	Verbose   bool
}

// NewListReposCommand creates a new `repos` command that lists the repos.
func NewListReposCommand(client hagen.Provider) *cobra.Command {
	var opts ListRepoOptions

	cmd := &cobra.Command{
		Use:   "repos",
		Short: "List the repositories based on a query. Default query is to list repos by ${GITHUB_OWNER}",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Args = args

			more := true
			query := ""
			received := 0
			for page := 1; more; page++ {
				query, searchOpts := NewSearchFromRepoOptions(opts)
				searchOpts.Page = page
				result, err := client.ListRepos(query, searchOpts)
				if err != nil {
					return err
				}

				err = DisplayRepos(result, os.Stdout)
				if err != nil {
					return err
				}

				received += len(result.Repositories)
				more = received < result.GetTotal()

				if !opts.Recursive {
					fmt.Printf("\nPress enter for more results\n")
					_, _ = fmt.Scanln()
				}
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
	flags.BoolVar(&opts.Recursive, "recursive", false, "Do you want to recursively get all results")
	flags.StringVar(&opts.Template, "template", "", "Use a query defined in the configuration file")
	flags.BoolVar(&opts.Verbose, "verbose", false, "Produce verbose output")

	return cmd
}

// NewSearchFromRepoOptions will return details that can perform a GitHub search
func NewSearchFromRepoOptions(opts ListRepoOptions) (string, github.SearchOptions) {
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
