package main

import (
	"flag"
	"fmt"
	"mcctl/controller"
	"os"
)

func main() {

	controller.CheckParams(os.Args)
	provider := os.Args[1]
	resource := os.Args[2]
	action := os.Args[3]

	if !controller.IsValidProvider(provider) {
		fmt.Println("Invalid provider")
		fmt.Println("\nAvailable providers: aws, azure, gcp")
		os.Exit(1)
	}
	
	controller.ExecuteAction(provider, resource, action)

	
}
