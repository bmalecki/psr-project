package main

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Reqeust events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

var svc *s3.S3
var uploader *s3manager.Uploader

func init() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		panic(err)
	}

	svc = s3.New(sess)
	uploader = s3manager.NewUploader(sess)
}

func createResponse(statusCode int, msg string) Response {
	return Response{
		StatusCode: 500,
		Body:       msg,
		Headers: map[string]string{
			"Content-Type": "plain/text",
		},
	}
}

func Handler(ctx context.Context, req Reqeust) (Response, error) {
	return createResponse(200, "OK"), nil
}

func main() {
	lambda.Start(Handler)
}
