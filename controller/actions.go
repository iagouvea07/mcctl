package controller

import (
	"fmt"
	"os"
)

func ExecuteActon(provider string, resource string, action string) {
	switch resource {
		case "instance":
			handleInstances(provider, action)
		default:
			fmt.Println("error")
	}
}

func handleInstances(provider string, action string) {
	switch action {

		case "delete":

		case "stop":

		case "reboot":

		case "ls":
			listInstances(provider)
		case "help":

		default:	
	}
}

func listInstances(provider string) {
	switch provider {
    	case "aws":
			fmt.Println("boa")
		case "azure":
			//do action
		case "gcp":
			//do action
		default:
			fmt.Println("error")
			os.Exit(1)
	}
}