package add

import (
	"fmt"
	"log"

	"github.com/reggles44/cs2-server-builder/pkg/lock"
	"github.com/reggles44/cs2-server-builder/pkg/plugins"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	command := cobra.Command{
		Use:   "add",
		Short: "Add a plugin to be used in building the zip",
		// Long:                   "List plugins that are supported either by available options or already installed options",
		// Example:                "list available",
		Run: func(cmd *cobra.Command, args []string) {
			l := lock.ReadLock()

			fmt.Println(l)
			for _, s := range args {
				plugin, err := plugins.FindPlugin(s)
				if err != nil {
					log.Fatal(err)
				} else {
					l.AddPlugin(plugin, plugin.Versions().Latest())
				}
			}

			l.WriteLock()
		},
	}

	return &command
}
