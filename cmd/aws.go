package cmd

import (
	"mcctl/controller"
	"github.com/spf13/cobra"
)

var awsOutput string

func init() {
	AwsCmd.Flags().StringVarP(&awsOutput, "output", "o", "table", "Output types (table, json)")
}

var AwsCmd = &cobra.Command{
	Use: "aws [resource] [action]",
	Short: "Option for interact with AWS resources",
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		controller.ExecuteAction(cmd.Name(), args[0], args[1], awsOutput)
	},	
}

