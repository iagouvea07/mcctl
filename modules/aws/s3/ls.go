package s3

import (
	"context"
	"fmt"
	"log"

	class "mcctl/classes"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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

	fmt.Println(bucketList)
}