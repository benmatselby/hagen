package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	hagen "github.com/benmatselby/hagen/pkg"
	"github.com/benmatselby/hagen/ui"
	"github.com/google/go-github/v72/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ListIssuesOptions defines what arguments/options the user can provide
type ListIssuesOptions struct {
	Args          []string
	Count         int
	DisplayLabels bool
	Query         string
	Recursive     bool
	Template      string
	Verbose       bool
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

				err = DisplayIssues(result, opts, os.Stdout)
				if err != nil {
					return err
				}

				received += len(result.Issues)
				more = received < result.GetTotal()

				if !opts.Recursive && more {
					fmt.Printf(ui.MoreResults)
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
	flags.IntVar(&opts.Count, "count", -1, "How many issues to display")
	flags.StringVar(&opts.Query, "query", "", "The search query to get issues")
	flags.BoolVar(&opts.Recursive, "recursive", false, "Do you want to recursively get all results")
	flags.StringVar(&opts.Template, "template", "", "Use a query defined in the configuration file")
	flags.BoolVar(&opts.Verbose, "verbose", false, "Produce verbose output")
	flags.BoolVar(&opts.DisplayLabels, "labels", false, "Whether we show labels")

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
func DisplayIssues(result *github.IssuesSearchResult, opts ListIssuesOptions, w io.Writer) error {
	for _, issue := range result.Issues {
		// Ultimately this should be pulled from something specific in the API,
		// but for now, just parse the API URL.
		repo := ""
		parts := strings.Split(issue.GetURL(), "/")
		if len(parts) > 5 {
			repo = fmt.Sprintf("%s/%s - ", parts[4], parts[5])
		}

		labels := ""
		if opts.DisplayLabels && len(issue.Labels) > 0 {
			labels += " ("
			for index, label := range issue.Labels {
				labels += label.GetName()
				if index < len(issue.Labels)-1 {
					labels += ", "
				}
			}
			labels += ")"
		}
		fmt.Fprintf(w, "- %s#%v %s%s\n", repo, issue.GetNumber(), issue.GetTitle(), labels)
	}
	return nil
}
