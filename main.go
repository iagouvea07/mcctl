package main

import (
	"fmt"
	"mcctl/cmd"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "mcctl",
	Short: "Multi Cloud Controller",
}

func init() {
	rootCmd.AddCommand(
		cmd.AwsCmd, 
		cmd.AzureCmd, 
		cmd.GcpCmd,
	)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err, os.Stderr)
		os.Exit(1)
	}
}
