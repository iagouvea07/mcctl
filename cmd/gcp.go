package cmd

import (
	"github.com/spf13/cobra"	
	"fmt"
)

var gcpOutput string

func init() {
	GcpCmd.Flags().StringVarP(&gcpOutput, "output", "o", "table", "Output types (table, json)")
}

var GcpCmd = &cobra.Command{
	Use: "gcp [resource] [action]",
	Short: "Option for interact with GCP resources",
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Name())
		fmt.Println(gcpOutput)
	},	
}

