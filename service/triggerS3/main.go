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

func pushImageKey(key string) (*sqs.SendMessageOutput, error) {
	result, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(0),
		MessageBody:  aws.String(key),
		QueueUrl:     &qURL,
	})

	return result, err
}

func Handler(ctx context.Context, s3Event events.S3Event) {
	for _, record := range s3Event.Records {
		s3 := record.S3
		_, err := pushImageKey(s3.Object.Key)

		if err != nil {
			fmt.Println("Error", err)
		}
	}
}

func main() {
	lambda.Start(Handler)
}
