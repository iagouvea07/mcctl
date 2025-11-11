package ec2

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/fatih/color"
)


type InstanceList struct {
	InstanceName string
	InstanceId string
	InstanceType string
	InstanceStatus string
	InstancePublicIp string
	InstancePrivateIp string
}

var (
	instanceId string
	instanceName string
	instanceType string
	instanceStatus string
	instancePublicIp string
	instancePrivateIp string
)

var jsonList []InstanceList

func ListInstances(output string) {

    cfg, err := config.LoadDefaultConfig(context.TODO(), 
        config.WithRegion("us-east-1"),
    )
    if err != nil {
        log.Fatal(err)
    }
	client := ec2.NewFromConfig(cfg)

	result, err := client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{})
    if err != nil {
        log.Fatal(err)
    }

	if result == nil {
		os.Exit(22)
	}

	for _, reservations := range result.Reservations {
		for _, instances := range reservations.Instances {

			for _, tag := range instances.Tags{if *tag.Key == "Name" {instanceName = *tag.Value}}
			if instanceName == "" {instanceName = "-"}
			if instances.PublicIpAddress != nil {instancePublicIp = *instances.PublicIpAddress } else {instancePublicIp = "-"}
			if instances.PrivateIpAddress != nil {instancePrivateIp = *instances.PrivateIpAddress} else {instancePrivateIp = "-"}
			instanceId = *instances.InstanceId
			instanceType = string(instances.InstanceType)
			instanceStatus = string(instances.State.Name)

			instance := InstanceList{
				InstanceName: instanceName,
				InstanceId: instanceId, 
				InstanceType: instanceType, 
				InstanceStatus: instanceStatus,
				InstancePublicIp: instancePublicIp,
				InstancePrivateIp: instancePrivateIp,
			}
			jsonList = append(jsonList, instance)
		}
	}
	jsonEncode, _ := json.MarshalIndent(jsonList, "", "  ")
	jsonList := string(jsonEncode)

	switch output {
		case "json": 
			fmt.Println(jsonList)
			os.Exit(0)

		case "table":
			tableOutput(jsonList)
	}
}

func tableOutput(result string) {
	var instances []InstanceList
	json.Unmarshal([]byte(result), &instances)

	headerColor := color.New(color.FgHiCyan, color.Bold)
	statusRunning := color.New(color.FgGreen, color.Bold)
	statusStopped := color.New(color.FgRed, color.Bold)
	statusOther := color.New(color.FgYellow)



	headerColor.Printf("%-20s %-25s %-12s %-15s %-20s %-20s\n",
		"NAME", "INSTANCE ID", "TYPE", "STATUS", "PUBLIC IP", "PRIVATE IP")

	for _, instance := range instances {
		var statusText string
		var statusSprintFunc func(a ...interface{}) string

		switch instance.InstanceStatus {
		case "running":
			statusText = "● running"
			statusSprintFunc = statusRunning.Sprint
		case "stopped":
			statusText = "■ stopped"
			statusSprintFunc = statusStopped.Sprint
		default:
			statusText = fmt.Sprintf("○ %s", instance.InstanceStatus)
			statusSprintFunc = statusOther.Sprint
		}

		name := fmt.Sprintf("%-20s", instance.InstanceName)
		instanceID := fmt.Sprintf("%-25s", instance.InstanceId)
		instanceType := fmt.Sprintf("%-12s", instance.InstanceType)
		status := fmt.Sprintf("%-15s", statusText)
		publicIP := fmt.Sprintf("%-20s", instance.InstancePublicIp)
		privateIP := fmt.Sprintf("%-20s", instance.InstancePrivateIp)

		fmt.Printf("%s %s %s %s %s %s\n",
			name,
			instanceID,
			instanceType,
			statusSprintFunc(status),
			publicIP,
			privateIP,
		)
	}
	fmt.Println()
}