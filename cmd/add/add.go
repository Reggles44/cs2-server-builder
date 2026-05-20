package add

import (
	"log"
	"os"

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
			l, err := lock.ReadLock()
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

			for _, s := range args {
				plugin := plugins.FindPlugin(s)
				if plugin != nil {
					l.AddPlugin(plugin plugins.Plugin, )
				} else {
					fmt.Println("%s is not a valid plugin", s)
				}
			}
		},
	}

	return &command
}
