package instance

import (
	"fmt"
	compute "mcctl/classes"
	"mcctl/modules/aws/ec2"
	"os"
)


func HandleInstances(provider string, action string, output string, instanceParam compute.InstanceParameters) {
	switch action {
		case "create":
			createInstances(provider, instanceParam)
		case "delete":
			deleteInstances(provider, instanceParam)
		case "stop":
			//TODO
		case "reboot":
			//TODO
		case "ls":
			listInstances(provider, output, instanceParam)
		case "help":
			//TODO
		default:	
			fmt.Println("error")
			os.Exit(1)
	}
}

func listInstances(provider string, output string, instanceParam compute.InstanceParameters) {
	switch provider {
    	case "aws":
			ec2.ListInstances(output, instanceParam)
		case "azure":
			//TODO
		case "gcp":
			//TODO
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
			//TODO
		case "gcp":
			//TODO
		default:
			fmt.Println("error")
			os.Exit(1)
	}
}

func deleteInstances(provider string, instanceParam compute.InstanceParameters) {
	switch provider {
    	case "aws":
			ec2.DeleteInstance(instanceParam)
		case "azure":
			//TODO
		case "gcp":
			//TODO
		default:
			fmt.Println("error")
			os.Exit(1)
	}	
}

