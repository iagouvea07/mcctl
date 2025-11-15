package bucket

import (
	"fmt"
	"mcctl/modules/aws/s3"
	"os"
)

func HandleBuckets(provider string, action string, output string){
	switch action {
		case "create":
			//TODO
		case "delete":
			//TODO

		case "ls":
			listBuckets(provider, output)
		case "help":

		default:	
			fmt.Println("error")
			os.Exit(1)		
	}
}

func listBuckets(provider string, output string){
	switch provider {
		case "aws":
			s3.ListBuckets(output)
		case "gcp":
			//TODO
		case "azure":
			//TODO
		default:
			fmt.Println("error")
			os.Exit(1)		
	}
}