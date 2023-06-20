package cmd

import (
    "os"
    "github.com/spf13/cobra"
)

var root_cmd = &cobra.Command{}

func Execute(){
    err := root_cmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}
