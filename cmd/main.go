package main

import (
	"log"
	"os"

	"github.com/reggles44/cs2-server-builder/cmd/add"
	"github.com/reggles44/cs2-server-builder/cmd/build"
	"github.com/reggles44/cs2-server-builder/cmd/list"
	"github.com/reggles44/cs2-server-builder/cmd/remove"
	"github.com/spf13/cobra"
)

func main() {
	command := cobra.Command{}

	command.AddCommand(add.NewCommand())
	command.AddCommand(build.NewCommand())
	command.AddCommand(list.NewCommand())
	command.AddCommand(remove.NewCommand())

	err := command.Execute()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
