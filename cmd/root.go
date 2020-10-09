package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//RootCmd is the entry point for the CLI
var RootCmd = &cobra.Command{
	Use:   "hosho",
	Short: "Hosho is a small language and server that makes policy enforcement easier using ABAC type rules.",
	Long:  "",
}

//Execute is
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr)
		os.Exit(1)
	}
}
