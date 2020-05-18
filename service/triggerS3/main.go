package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Reqeust events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

// var svc *s3.S3
// var uploader *s3manager.Uploader

// func init() {
// 	sess, err := session.NewSession(&aws.Config{
// 		Region: aws.String("us-east-1")},
// 	)
// 	if err != nil {
// 		panic(err)
// 	}

// 	svc = s3.New(sess)
// 	uploader = s3manager.NewUploader(sess)
// }

func Handler(ctx context.Context, s3Event events.S3Event) {
	for _, record := range s3Event.Records {
		s3 := record.S3
		fmt.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3.Bucket.Name, s3.Object.Key)
	}
}

func main() {
	lambda.Start(Handler)
}
