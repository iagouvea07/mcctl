package controllers

import (
	"fmt"
	class "mcctl/classes"
	"mcctl/controllers/bucket"
	"mcctl/controllers/instance"
)

func ExecuteAction(
		provider string, 
		resource string, 
		action string, 
		output string, 
		instanceParam class.InstanceParameters,
		bucketParam class.BucketParameters,
	) {
	switch resource {
		case "instance":
			instance.HandleInstances(provider, action, output, instanceParam)
		case "bucket":
			bucket.HandleBuckets(provider, action, output, bucketParam)
		default:
			fmt.Println("error")
	}
}
