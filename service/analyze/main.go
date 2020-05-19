package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/textract"
	"github.com/psr-project/uploadService/imageservice"
)

var imageTableSvc *imageservice.ImageTableService
var textractSvc *textract.Textract
var sqsSvc *sqs.SQS
var analyzeImageQUrl string
var textractQUrl string

func init() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	))

	imageTableSvc = imageservice.New(sess, os.Getenv("IMAGE_TABLE"))
	textractSvc = textract.New(sess)
	sqsSvc = sqs.New(sess)

	analyzeImageQUrl = os.Getenv("ANALYZE_IMAGE_QUEUE_URL")
	textractQUrl = os.Getenv("TEXTRACT_QUEUE_URL")
}

func Handler(ctx context.Context, sqsEvent events.SQSEvent) error {

	for _, record := range sqsEvent.Records {
		fileName := record.Body
		if err := imageTableSvc.ProcessingImageTableItem(fileName); err != nil {
			return err
		}

		if _, err := sqsSvc.DeleteMessage(&sqs.DeleteMessageInput{
			QueueUrl:      &analyzeImageQUrl,
			ReceiptHandle: &record.ReceiptHandle,
		}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lambda.Start(Handler)
}
