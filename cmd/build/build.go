package build

import "github.com/spf13/cobra"

func NewCommand() *cobra.Command {
	command := cobra.Command{
		Use:   "build",
		Short: "Build the zip ",
		Run: func(cmd *cobra.Command, args []string) {
			panic("TODO")
		},
	}

	return &command
}
