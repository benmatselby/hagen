package repo

import (
	hagen "github.com/benmatselby/hagen/pkg"
	"github.com/spf13/cobra"
)

// NewRepoCommand creates a new `repo` command
func NewRepoCommand(client hagen.Provider) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repo",
		Short: "Repository related commands",
	}

	cmd.AddCommand(
		NewRepoLsCommand(client),
	)
	return cmd
}
