package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewListTemplatesCommand creates the `templates` command
func NewListTemplatesCommand() *cobra.Command {
	var verbose bool

	cmd := &cobra.Command{
		Use:   "templates",
		Short: "List all templates from the config file",
		Run: func(cmd *cobra.Command, args []string) {
			templates := viper.GetStringMap("templates")
			w := cmd.OutOrStdout()
			if len(templates) == 0 {
				fmt.Fprintln(w, "No templates found in config file.")
				return
			}

			// Sort template names for consistent output
			names := make([]string, 0, len(templates))
			for name := range templates {
				names = append(names, name)
			}
			sort.Strings(names)

			// Extract all template details once for efficiency
			templateDetails := make(map[string]map[string]interface{}, len(names))
			if verbose {
				for _, name := range names {
					if tmpl, ok := templates[name].(map[string]interface{}); ok {
						templateDetails[name] = tmpl
					} else {
						templateDetails[name] = viper.GetStringMap(fmt.Sprintf("templates.%s", name))
					}
				}
			}

			for _, name := range names {
				fmt.Fprint(w, name)
				if verbose {
					tmpl := templateDetails[name]
					if query, ok := tmpl["query"].(string); ok {
						fmt.Fprintf(w, ": %s", query)
					}
				}
				fmt.Fprintln(w, "")
			}
		},
	}

	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Show the query for each template")
	return cmd
}
