package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/psr-project/uploadService/imageservice"
)

var imageTableService *imageservice.ImageTableService

func init() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		panic(err)
	}

	imageTableService = imageservice.New(sess, os.Getenv("IMAGE_TABLE"))
}

func Handler(ctx context.Context, sqsEvent events.SQSEvent) error {

	for _, record := range sqsEvent.Records {
		fileName := record.Body
		err := imageTableService.ProcessingImageTableItem(fileName)
		fmt.Print(err.Error())
	}
	return nil
}

func main() {
	lambda.Start(Handler)
}
