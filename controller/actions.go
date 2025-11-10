package controller

import (
	"fmt"
	compute "mcctl/classes"
	"mcctl/modules/ec2"
	"os"
)

func ExecuteAction(provider string, resource string, action string, output string, instanceParam compute.InstanceParameters) {
	switch resource {
		case "instance":
			handleInstances(provider, action, output, instanceParam)
		default:
			fmt.Println("error")
	}
}

func handleInstances(provider string, action string, output string, instanceParam compute.InstanceParameters) {
	switch action {
		case "create":
			createInstances(provider, instanceParam)
		case "delete":

		case "stop":

		case "reboot":

		case "ls":
			listInstances(provider, output)
		case "help":

		default:	
	}
}

func listInstances(provider string, output string) {
	switch provider {
    	case "aws":
			ec2.ListInstances(output)
		case "azure":
			//do action
		case "gcp":
			//do action
		default:
			fmt.Println("error")
			os.Exit(1)
	}
}

func createInstances(provider string, instanceParam compute.InstanceParameters) {
	switch provider {
    	case "aws":
			ec2.CreateInstance(instanceParam)
		case "azure":
			//do action
		case "gcp":
			//do action
		default:
			fmt.Println("error")
			os.Exit(1)
	}
}