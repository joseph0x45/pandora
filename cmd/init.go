package cmd

import (
	"github.com/spf13/cobra"
)

var init_cmd = &cobra.Command{
    Use: "init",
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func init(){
    root_cmd.AddCommand(init_cmd)
}
