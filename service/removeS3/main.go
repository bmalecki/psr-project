package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/psr-project/uploadService/imageservice"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Reqeust events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

var svcS3 *s3.S3
var imageTableService *imageservice.ImageTableService
var tableName string
var bucketId string

func init() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	))

	svcS3 = s3.New(sess)
	imageTableService = imageservice.New(sess, os.Getenv("IMAGE_TABLE"))
	bucketId = os.Getenv("UPLOAD_IMAGE_STORAGE_ID")
}

func createResponse(statusCode int, msg string) Response {
	return Response{
		StatusCode: statusCode,
		Body:       msg,
		Headers: map[string]string{
			"Content-Type":                     "application/json",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}
}

func Handler(ctx context.Context, req Reqeust) (Response, error) {
	id := req.PathParameters["id"]

	err := imageTableService.DeleteImageItemById(id)

	_, err = svcS3.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(bucketId), Key: aws.String(id)})
	if err != nil {
		return createResponse(500, ""), fmt.Errorf("Unable to delete object %q from bucket %q, %v", id, bucketId, err)
	}

	err = svcS3.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucketId),
		Key:    aws.String(id),
	})

	if err != nil {
		fmt.Println(err.Error())
		return createResponse(500, ""), err
	}

	return createResponse(200, id), nil
}

func main() {
	lambda.Start(Handler)
}
