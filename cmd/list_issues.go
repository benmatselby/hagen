package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	hagen "github.com/benmatselby/hagen/pkg"
	"github.com/benmatselby/hagen/ui"
	"github.com/google/go-github/v72/github"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ListIssuesOptions defines what arguments/options the user can provide
type ListIssuesOptions struct {
	Args            []string
	Count           int
	DisplayLabels   bool
	Query           string
	Recursive       bool
	Template        string
	Verbose         bool
	DisplayStrategy string // new flag for strategy
	HumanDates      bool   // new flag for human-readable dates
}

// NewListIssuesCommand will register the `issues` command
func NewListIssuesCommand(client hagen.Provider) *cobra.Command {
	var opts ListIssuesOptions

	cmd := &cobra.Command{
		Use:   "issues",
		Short: "List issues given the search criteria. Default query is to list issues where the author is ${GITHUB_OWNER}",
		Long: `This command will return issues and pull requests given the search criteria.
`,
		Example: `
Get all issues/pull requests raised by a user:

$ hagen issues list --query "author:yourusername"

Get all pull requests reviewed by a user:

$ hagen issues list --query "reviewed-by:yourusername"

Get all issues/pull requests that a user is involved with:

$ hagen issues list --query "involves:yourusername"

Display the results in a table format:

$ hagen issues list --query "author:yourusername" --display table

Display the results in a human-readable format:

$ hagen issues list --query "author:yourusername" --display table --human-dates
`,
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

				// Select strategy
				var strategy IssueDisplayStrategy
				switch opts.DisplayStrategy {
				case "table":
					strategy = TableIssueDisplayStrategy{}
				default:
					strategy = DefaultIssueDisplayStrategy{}
				}

				err = strategy.Display(result, opts, os.Stdout)
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
	flags.StringVar(&opts.DisplayStrategy, "display", "default", "Display strategy: default or table")
	flags.BoolVar(&opts.HumanDates, "human-dates", false, "Display dates in a human-readable format")

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
	} else {
		if opts.Query != "" {
			query = opts.Query
		}

		if opts.Count != -1 {
			searchOpts.PerPage = opts.Count
		}
	}

	return query, searchOpts
}

// IssueDisplayStrategy defines the interface for displaying issues
// You can add more strategies by implementing this interface
type IssueDisplayStrategy interface {
	Display(result *github.IssuesSearchResult, opts ListIssuesOptions, w io.Writer) error
}

// DefaultIssueDisplayStrategy implements the default display logic
// (current behaviour)
type DefaultIssueDisplayStrategy struct{}

func (s DefaultIssueDisplayStrategy) Display(result *github.IssuesSearchResult, opts ListIssuesOptions, w io.Writer) error {
	for _, issue := range result.Issues {
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

// TableIssueDisplayStrategy implements table display logic
type TableIssueDisplayStrategy struct{}

func (s TableIssueDisplayStrategy) Display(result *github.IssuesSearchResult, opts ListIssuesOptions, w io.Writer) error {
	table := tablewriter.NewWriter(w)
	table.Header([]string{"Type", "Repository", "Number", "Title", "Labels", "Status", "Created At", "Closed At"})

	for _, issue := range result.Issues {
		repo := ""
		parts := strings.Split(issue.GetURL(), "/")
		if len(parts) > 5 {
			repo = fmt.Sprintf("%s/%s", parts[4], parts[5])
		}
		labels := ""
		if opts.DisplayLabels && len(issue.Labels) > 0 {
			for index, label := range issue.Labels {
				labels += label.GetName()
				if index < len(issue.Labels)-1 {
					labels += ", "
				}
			}
		}
		typeStr := "Issue"
		if issue.IsPullRequest() {
			typeStr = "Pull Request"
		}
		status := ""
		if issue.State != nil {
			status = *issue.State
		}
		createdAt := ""
		if issue.CreatedAt != nil {
			if opts.HumanDates {
				createdAt = issue.CreatedAt.Format(ui.HumanFriendlyDateFormat)
			} else {
				createdAt = issue.CreatedAt.Format(ui.DateFormat)
			}
		}
		closedAt := ""
		if issue.ClosedAt != nil {
			if opts.HumanDates {
				closedAt = issue.ClosedAt.Format(ui.HumanFriendlyDateFormat)
			} else {
				closedAt = issue.ClosedAt.Format(ui.DateFormat)
			}
		}
		row := []string{typeStr, repo, fmt.Sprintf("%v", issue.GetNumber()), issue.GetTitle(), labels, status, createdAt, closedAt}
		err := table.Append(row)
		if err != nil {
			return err
		}
	}
	return table.Render()
}

// DisplayIssues is deprecated. Use IssueDisplayStrategy instead.
// This is kept for backward compatibility and for tests.
func DisplayIssues(result *github.IssuesSearchResult, opts ListIssuesOptions, w io.Writer) error {
	return DefaultIssueDisplayStrategy{}.Display(result, opts, w)
}
