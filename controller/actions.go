package controller

import (
	"fmt"
	"mcctl/modules/ec2"
	"os"
)

func ExecuteAction(provider string, resource string, action string, output string) {
	switch resource {
		case "instance":
			handleInstances(provider, action, output)
		default:
			fmt.Println("error")
	}
}

func handleInstances(provider string, action string, output string) {
	switch action {

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