package issue

import (
	hagen "github.com/benmatselby/hagen/pkg"
	"github.com/spf13/cobra"
)

// NewIssueCommand creates a new `issue` command
func NewIssueCommand(client hagen.Provider) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issue",
		Short: "Issue related commands",
	}

	cmd.AddCommand(
		NewLsCommand(client),
	)
	return cmd
}
