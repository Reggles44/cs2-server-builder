package remove

import "github.com/spf13/cobra"

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:     "remove",
		Short:   "Remove a plugin",
		Example: "remove sharptimer",
		Run: func(cmd *cobra.Command, args []string) {
			panic("TODO")
		},
	}

	return command
}
