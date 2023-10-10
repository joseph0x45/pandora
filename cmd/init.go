package cmd

import (
	"github.com/spf13/cobra"
)

var (
  init_message =
`Welcome to Pandora, a simple and easy to use password manager.
You are creating a new Pandora box. Enter the key to the box:`
)

var init_cmd = &cobra.Command{
	Use: "init",
	Run: func(cmd *cobra.Command, args []string) {
    println(init_message)

	},
}

func init() {
	root_cmd.AddCommand(init_cmd)
}
