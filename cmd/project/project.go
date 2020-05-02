package project

import (
	hagen "github.com/benmatselby/hagen/pkg"
	"github.com/spf13/cobra"
)

// NewProjectCommand creates a new `project` command
func NewProjectCommand(client hagen.Provider) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project",
		Short: "Project related commands",
	}

	cmd.AddCommand(
		NewLsCommand(client),
		NewBurndownCommand(client),
	)
	return cmd
}
