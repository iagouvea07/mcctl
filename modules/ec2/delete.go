package ec2

import (
	"context"
	"fmt"
	"log"
	"mcctl/classes"
	"os"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func DeleteInstance(parameters compute.InstanceParameters) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), 
		config.WithRegion("us-east-1"),
	)

	if err != nil {
		log.Fatal(err)
	}

	client := ec2.NewFromConfig(cfg)

	var instanceList []string
	if parameters.InstanceName != "" {
		searchEesult, searchErr := client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{
			Filters: []types.Filter{
				{
					Name: aws.String("tag:Name"),
					Values: []string{parameters.InstanceName}, 
				},
			},
		})

		if searchErr != nil {
			log.Fatal(searchErr)
		}

		for _, reservations := range searchEesult.Reservations {
			for _, instances := range reservations.Instances {
				for _, tag := range instances.Tags {
					if *tag.Key == "Name" && (instances.State.Name == "running" || instances.State.Name == "pending") {
						instanceId = *instances.InstanceId
						parameters.InstanceId = instanceId
						instanceList = append(instanceList, instanceId)
					}
				}
			}
		}

		_, DeleteErr := client.TerminateInstances(context.TODO(), &ec2.TerminateInstancesInput{
			InstanceIds: instanceList,
		})

		if DeleteErr != nil {
			fmt.Println("The InstanceId is required [--instance]")
			os.Exit(1)
		}

		fmt.Println(instanceList, " is Deleted!")
		
	} else {
		_, DeleteErr := client.TerminateInstances(context.TODO(), &ec2.TerminateInstancesInput{
			InstanceIds: []string{parameters.InstanceId},
		})

		if DeleteErr != nil {
			fmt.Println("The InstanceId is required [--instance]")
			os.Exit(1)
		}

		fmt.Println("[",parameters.InstanceId,"]", " is Deleted!")
	}

}
		