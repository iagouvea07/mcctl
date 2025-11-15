package controllers

import (
	"fmt"
	compute "mcctl/classes"
	"mcctl/controllers/bucket"
	"mcctl/controllers/instance"
)

func ExecuteAction(
		provider string, 
		resource string, 
		action string, 
		output string, 
		instanceParam compute.InstanceParameters,
	) {
	switch resource {
		case "instance":
			instance.HandleInstances(provider, action, output, instanceParam)
		case "bucket":
			bucket.HandleBuckets(provider, action, output)
		default:
			fmt.Println("error")
	}
}
