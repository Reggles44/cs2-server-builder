package list

import (
	"fmt"

	"github.com/reggles44/cs2-server-builder/pkg/plugins"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:     "list",
		Short:   "List plugins that are available",
		Long:    "List plugins, their respective repos, and available versions",
		Example: "list",
		Run: func(cmd *cobra.Command, args []string) {
			for _, plugin := range plugins.Map {
				fmt.Println(plugin.String())
			}
		},
	}

	return command
}
