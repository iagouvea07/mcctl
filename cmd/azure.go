package cmd

import (
	"github.com/spf13/cobra"	
	"fmt"
)

var azureOutput string

func init() {
	AzureCmd.Flags().StringVarP(&azureOutput, "output", "o", "table", "Output types (table, json)")
}

var AzureCmd = &cobra.Command{
	Use: "azure [resource] [action]",
	Short: "Option for interact with Azure resources",
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Name())
		fmt.Println(azureOutput)
	},	
}

