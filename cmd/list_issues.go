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

// ListIssuesOptions defines what arguments/options the user can provide
type ListIssuesOptions struct {
	Args      []string
	Count     int
	Query     string
	Recursive bool
	Template  string
	Verbose   bool
}

// NewListIssuesCommand will register the `issues` command
func NewListIssuesCommand(client hagen.Provider) *cobra.Command {
	var opts ListIssuesOptions

	cmd := &cobra.Command{
		Use:   "issues",
		Short: "List issues given the search criteria. Default query is to list issues where the author is ${GITHUB_OWNER}",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Args = args

			more := true
			query := ""
			received := 0
			for page := 1; more; page++ {
				query, searchOpts := NewSearchFromIssueOptions(opts)
				searchOpts.Page = page
				result, err := client.ListIssues(query, searchOpts)
				if err != nil {
					return err
				}

				err = DisplayIssues(result, os.Stdout)
				if err != nil {
					return err
				}

				received += len(result.Issues)
				more = received < result.GetTotal()

				if !opts.Recursive {
					fmt.Printf("\nPress enter for more results\n")
					fmt.Scanln()
				}
			}

			if opts.Verbose {
				fmt.Fprintf(os.Stdout, "\n\nQuery used: %s\n", query)
			}

			return nil
		},
	}

	flags := cmd.Flags()
	flags.IntVar(&opts.Count, "count", -1, "How many issues to display")
	flags.StringVar(&opts.Query, "query", "", "The search query to get issues")
	flags.BoolVar(&opts.Recursive, "recursive", false, "Do you want to recursively get all results")
	flags.StringVar(&opts.Template, "template", "", "Use a query defined in the configuration file")
	flags.BoolVar(&opts.Verbose, "verbose", false, "Produce verbose output")

	return cmd
}

// NewSearchFromIssueOptions will return details that can perform a GitHub search
func NewSearchFromIssueOptions(opts ListIssuesOptions) (string, github.SearchOptions) {
	query := fmt.Sprintf("author:%s", viper.GetString("GITHUB_OWNER"))
	query += fmt.Sprintf(" state:%s", "open")
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

		if opts.Count != -1 {
			searchOpts.ListOptions.PerPage = opts.Count
		}
	}

	return query, searchOpts
}

// DisplayIssues will display issues based on the given search criteria
func DisplayIssues(result *github.IssuesSearchResult, w io.Writer) error {
	for _, issue := range result.Issues {
		fmt.Fprintf(w, "- #%v %s\n", issue.GetNumber(), issue.GetTitle())
	}
	return nil
}
