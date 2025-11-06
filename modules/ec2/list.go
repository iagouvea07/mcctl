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


func ListInstances(output string) {
	var instanceId string
	var instanceName string
	var instanceType string
	var instanceStatus string
	var instancePublicIp string
	var instancePrivateIp string

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

	w := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', tabwriter.Debug)

	blue := color.New(color.FgCyan).Add(color.Underline)
	white := color.New(color.FgWhite).Add(color.Underline)

	if output == "json" {
		fmt.Println("saida", output)

		type InstanceList struct {
			InstanceId string
			InstanceType string
			instanceStatus string
		}

		var jsonList []InstanceList

		for _, reservations := range result.Reservations {
			for _, instances := range reservations.Instances {
				instance := InstanceList{
					InstanceId: *instances.InstanceId, 
					InstanceType: string(instances.InstanceType), 
					instanceStatus: string(instances.State.Name),
				}
				jsonList = append(jsonList, instance)
			}
		}
		jsonEncode, _ := json.Marshal(jsonList)
		fmt.Println(string(jsonEncode))
	} else {
		blue.Fprintf(w, "%-20s %-25s %-15s %-10s %-15s %-15s\n",
		"NAME", "INSTANCE ID", "TYPE", "STATUS", "PUBLIC IP", "PRIVATE IP")
		for _, reservations := range result.Reservations {
			for _, instance := range reservations.Instances {
				for _, tag := range instance.Tags{if *tag.Key == "Name" {instanceName = *tag.Value}}
				instanceId = *instance.InstanceId
				instanceType = string(instance.InstanceType)
				instanceStatus = string(instance.State.Name)

				if instance.PublicIpAddress != nil {instancePublicIp = *instance.PublicIpAddress } else {instancePublicIp = "-"}
				if instance.PrivateIpAddress != nil {instancePrivateIp = *instance.PrivateIpAddress} else {instancePrivateIp = "-"}
				
				white.Fprintf(w, "%-20s %-25s %-15s %-10s %-15s %-15s\n",
					instanceName, instanceId, instanceType, instanceStatus, instancePublicIp, instancePrivateIp)
				w.Flush()
			}
		}
	}
	
}
