package cmd

import (
	"mcctl/classes"
	"mcctl/controllers"
	"github.com/spf13/cobra"
)


var awsOutput string
var instanceParameters class.InstanceParameters

func init() {

	//AWS flags
	AwsCmd.Flags().StringVarP(&awsOutput, "output", "o", "table", "Output types (table, json)")
	AwsCmd.Flags().StringVar(&instanceParameters.InstanceName, "name", "", "Define an instance name")
	AwsCmd.Flags().StringVar(&instanceParameters.InstanceAmi, "image", "ami-084568db4383264d4", "Select the instance AMI")
	AwsCmd.Flags().StringVar(&instanceParameters.InstanceType, "type", "t3.micro", "Select the instance size")
	AwsCmd.Flags().StringVar(&instanceParameters.InstanceKey, "key", "", "Select your key pair (required)")
	AwsCmd.Flags().StringVar(&instanceParameters.InstanceId, "instance", "", "Select your instance Id")

}

var AwsCmd = &cobra.Command{
	Use: "aws [resource] [action]",
	Short: "Option for interact with AWS resources",
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		controllers.ExecuteAction(cmd.Name(), args[0], args[1], awsOutput, instanceParameters)
	},	
}