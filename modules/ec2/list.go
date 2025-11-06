package ec2

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
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

	w := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', tabwriter.Debug)
	blue := color.New(color.FgCyan).Add(color.Underline)
	white := color.New(color.FgWhite).Add(color.Underline)

	blue.Fprintf(w, "%-30s %-25s %-15s %-10s %-15s %-15s\n",
		"NAME", "INSTANCE ID", "TYPE", "STATUS", "PUBLIC IP", "PRIVATE IP")

	for _, instance := range instances {
		white.Fprintf(w, "%-30s %-25s %-15s %-10s %-15s %-15s\n",
			instance.InstanceName,
			instance.InstanceId,
			instance.InstanceType,
			instance.InstanceStatus,
			instance.InstancePublicIp,
			instance.InstancePrivateIp,
		)
		w.Flush()
	}
}