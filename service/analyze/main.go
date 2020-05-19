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

func Handler(ctx context.Context, req Reqeust) error {
	fmt.Println("Recieved")
	return nil
}

func main() {
	lambda.Start(Handler)
}
