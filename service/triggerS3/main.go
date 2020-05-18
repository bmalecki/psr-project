package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Reqeust events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

var svc *sqs.SQS
var qURL string

func init() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		panic(err)
	}

	svc = sqs.New(sess)
	qURL = os.Getenv("ANALYZE_IMAGE_QUEUE_URL")
}

func pushRecord(qURL string) (*sqs.SendMessageOutput, error) {

	result, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		// MessageAttributes: map[string]*sqs.MessageAttributeValue{
		// 	"Title": {
		// 		DataType:    aws.String("String"),
		// 		StringValue: aws.String("The Whistler"),
		// 	},
		// 	"Author": {
		// 		DataType:    aws.String("String"),
		// 		StringValue: aws.String("John Grisham"),
		// 	},
		// 	"WeeksOn": {
		// 		DataType:    aws.String("Number"),
		// 		StringValue: aws.String("6"),
		// 	},
		// },
		MessageBody: aws.String("Information about current NY Times fiction bestseller for week of 12/11/2016."),
		QueueUrl:    &qURL,
	})

	return result, err
}

func Handler(ctx context.Context, s3Event events.S3Event) {
	for _, record := range s3Event.Records {
		s3 := record.S3
		fmt.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3.Bucket.Name, s3.Object.Key)
	}

	result, err := pushRecord(qURL)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Success", *result.MessageId)
}

func main() {
	lambda.Start(Handler)
}
