package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/psr-project/uploadService/imageservice"
)

var imageTableSvc *imageservice.ImageTableService
var sqsSvc *sqs.SQS
var textractImageQUrl string

// func init() {
// sess := session.Must(session.NewSession(&aws.Config{
// 	Region: aws.String("us-east-1")},
// ))

// imageTableSvc = imageservice.New(sess, os.Getenv("IMAGE_TABLE"))
// sqsSvc = sqs.New(sess)

// textractImageQUrl = os.Getenv("TEXTRACT_QUEUE_URL")
// }

func Handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {
		fmt.Printf("The message %s for event source %s = %s \n", message.MessageId, message.EventSource, message.Body)
	}

	return nil
}

// func Handler(ctx context.Context, sqsEvent events.SQSEvent) error {
// fmt.Println("Start")

// for _, record := range sqsEvent.Records {
// 	body := record.Body

// 	fmt.Printf("Body: %v", body)

// if err := imageTableSvc.ReadyImageStatusItem(fileName); err != nil {
// 	return err
// }

// if _, err := sqsSvc.DeleteMessage(&sqs.DeleteMessageInput{
// 	QueueUrl:      &textractImageQUrl,
// 	ReceiptHandle: &record.ReceiptHandle,
// }); err != nil {
// 	return err
// }

// }
// return nil
// }

func main() {
	lambda.Start(Handler)
}
