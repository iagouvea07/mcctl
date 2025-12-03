package bucket

import (
	"fmt"
	class "mcctl/classes"
	"mcctl/modules/aws/s3"
	"os"
)

func HandleBuckets(provider string, action string, output string, bucketParam class.BucketParameters){
	switch action {
		case "create":
			//TODO
		case "delete":
			//TODO

		case "ls":
			listBuckets(provider, output)
		case "help":

		case "copy":
			copyObject(provider, output, bucketParam)
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

func copyObject(provider string, output string, bucketParam class.BucketParameters) {
	switch provider {
		case "aws":
			s3.CopyFiles(bucketParam, output)
		case "gcp":
			//TODO
		case "azure":
			//TODO
		default:
			fmt.Println("error")
			os.Exit(1)		
	}	
}