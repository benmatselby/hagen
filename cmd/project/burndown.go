package project

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/tabwriter"
	"github.com/spf13/cobra"
	hagen "github.com/benmatselby/hagen/pkg"
)

// BurndownOptions defines what arguments/options the user can provide for the
// `project burndown` command.
type BurndownOptions struct {
	Args    []string
	Org string
	Project string
	Repo    string
}

// NewBurndownCommand creates a new `project burndown` command that provides
// tabular data for the given project.
func NewBurndownCommand(client hagen.Provider) *cobra.Command {
	var opts BurndownOptions

	cmd := &cobra.Command{
		Use:   "burndown",
		Short: "Provide a burndown table for a project",
		Long:  "Requires prefixes of '(x) ' in front of all issue titles. The x is the estimation value.",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Args = args
			return DisplayBurndown(client, opts, os.Stdout)
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&opts.Project, "project", "", "The project to produce burndown for")
	flags.StringVar(&opts.Org, "org", "", "The organisation the project is related to")
	flags.StringVar(&opts.Repo, "repo", "", "The repo the project is related to")

	return cmd
}

// DisplayBurndown will display tabular data for the given project
func DisplayBurndown(client hagen.Provider, userOpts BurndownOptions, w io.Writer) error {
	if userOpts.Project == "" {
		return fmt.Errorf("a project needs to be specified --project")
	}

	columns, err := client.ListColumnsForProject(userOpts.Project, userOpts.Org, userOpts.Repo)
	if err != nil {
		return err
	}

	tw := tabwriter.NewWriter(w, 0, 1, 1, ' ', 0)
	fmt.Fprintf(tw, "%s\t%s\t%s\n", "Column", "Cards", "Story Points")
	fmt.Fprintf(tw, "%s\t%s\t%s\n", "------", "-----", "------------")

	totalCard := 0
	totalPoint := 0
	for _, column := range columns {
		cardCount := 0
		pointCount := 0
		issues, err := client.ListIssuesForProjectColumn(column.GetID())
		if err != nil {
			return err
		}

		for _, issue := range issues {
			re := regexp.MustCompile(`^\(.+?\)`)
			title := re.FindString(issue.GetTitle())
			points, err := strconv.Atoi(strings.Trim(title, "()"))
			if err != nil {
				points = 0
			}
			pointCount += points
			cardCount++
		}
		totalCard += cardCount
		totalPoint += pointCount
		fmt.Fprintf(tw, "%s\t%v\t%v\n", column.GetName(), cardCount, pointCount)
	}

	fmt.Fprintf(tw, "%s\t%s\t%s\n", "-----", "-----", "------------")
	fmt.Fprintf(tw, "%s\t%d\t%d\n", "Total", totalCard, totalPoint)
	fmt.Fprintf(tw, "%s\t%s\t%s\n", "-----", "-----", "------------")

	tw.Flush()

	return nil
}
