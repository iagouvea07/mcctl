package s3

import (
	"context"
	"fmt"
	"log"
	class "mcctl/classes"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func CopyFiles(param class.BucketParameters, output string) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))

	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	bucket := &param.BucketName
	object := &param.FileName
	content, _ := os.Open(*object)

	result, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(*bucket),
		Key: aws.String(*object),
		Body: content,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}