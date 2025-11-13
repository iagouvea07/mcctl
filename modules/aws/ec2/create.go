package ec2

import (
	"context"
	"fmt"
	"log"
	"os"
	"mcctl/classes"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func CreateInstance(parameters compute.InstanceParameters) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
	)

	if err != nil {
		log.Fatal(err)
	}


	client := ec2.NewFromConfig(cfg)

	if parameters.InstanceKey == "" { 
		fmt.Println("An key pair is required to create an instance [--key key-name]")
		os.Exit(1)
	}

	result, err := client.RunInstances(context.TODO(), &ec2.RunInstancesInput{
		ImageId: aws.String(parameters.InstanceAmi),
		InstanceType: types.InstanceType(parameters.InstanceType),
		KeyName: aws.String(parameters.InstanceKey),
		MinCount: aws.Int32(1),
		MaxCount: aws.Int32(1),
	})

	if err != nil {
		fmt.Println("Could not create instance: ", err)
        return
	}

	instanceId := result.Instances[0].InstanceId

	_, tagErr := client.CreateTags(context.TODO(), &ec2.CreateTagsInput{
		Resources: []string{
			*instanceId, 
		},
		Tags: []types.Tag{
			{
				Key: aws.String("Name"),
				Value: aws.String(parameters.InstanceName),
			},
		},
	})

	if tagErr != nil {
		fmt.Printf("Could not create tags: %v\n", tagErr)
		return
	}

	fmt.Printf("🚀 Instance Created: %s\n\n", *instanceId)
}