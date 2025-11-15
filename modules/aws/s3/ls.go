package s3

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	class "mcctl/classes"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/fatih/color"
)

func ListBuckets(output string) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))

	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	DescribeResult, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	var bucketList []class.BucketDescribe

	if err != nil {
		log.Fatal(err)
	}

	for _, bucket := range DescribeResult.Buckets {

		bucketInfo := class.BucketDescribe{
			BucketName: *bucket.Name,
		}
		bucketList = append(bucketList, bucketInfo)
	}

	jsonEncode, _ := json.MarshalIndent(bucketList, "", "  ")

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
	var buckets []class.BucketDescribe
	json.Unmarshal([]byte(result), &buckets)

	headerColor := color.New(color.FgHiCyan, color.Bold)

	headerColor.Printf("%-25s\n", "BUCKET NAME")

	for _, bucket := range buckets {
		name := fmt.Sprintf("%-25s", bucket.BucketName)

		fmt.Printf("%s\n", name)
	}
}