package controller

import (
		"fmt" 
		"os"
	)

func CheckParams(args []string) {
	if len(args) < 2 {
		fmt.Println("Need provider")
		fmt.Println("Available Providers: aws, azure, gcp")
		os.Exit(1)
	} 

	if len(args) < 3 {
		fmt.Println("Need resource")
		fmt.Println("Available Resources: instance, loadbalancer, database, cdn")
		os.Exit(1)
	}

	if len(args) < 4 {
		fmt.Println("Need Action")
		fmt.Println("Use **help** action to list more options for this resource")
		os.Exit(1)
	}
}

func IsValidProvider(provider string) bool {
	providers := []string{"aws", "azure", "gcp"}
	for _, p := range providers {
		if p == provider {return true }
	}
	return false
}
