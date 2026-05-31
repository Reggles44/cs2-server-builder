package remove

import (
	"log"

	"github.com/reggles44/cs2-server-builder/pkg/lock"
	"github.com/reggles44/cs2-server-builder/pkg/plugins"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:     "remove",
		Short:   "Remove a plugin",
		Example: "remove sharptimer",
		Run: func(cmd *cobra.Command, args []string) {
			l := lock.ReadLock()

			for _, s := range args {
				plugin, err := plugins.FindPlugin(s)
				if err != nil {
					log.Fatal(err)
				} else {
					l.RemovePlugin(plugin)
				}
			}

			l.WriteLock()
		},
	}

	return command
}
