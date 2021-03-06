package main

import (
	"context"
	"encoding/json"
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

func init() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	))

	svcS3 = s3.New(sess)
	imageTableService = imageservice.New(sess, os.Getenv("IMAGE_TABLE"))
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
	var imageItemJson []byte
	var err error
	id := req.PathParameters["id"]

	fmt.Println(id)

	if len(id) != 0 {
		imageItem, err := imageTableService.GetImageItemById(id)

		if err != nil {
			fmt.Println(err.Error())
			return createResponse(500, ""), err
		}

		imageItemJson, err = json.Marshal(imageItem)
	} else {
		imageItemList, err := imageTableService.GetAllImageItems()

		if err != nil {
			fmt.Println(err.Error())
			return createResponse(500, ""), err
		}

		imageItemJson, err = json.Marshal(imageItemList)
	}

	if err != nil {
		fmt.Println(err.Error())
		return createResponse(500, ""), err
	}

	return createResponse(200, string(imageItemJson)), nil
}

func main() {
	lambda.Start(Handler)
}
