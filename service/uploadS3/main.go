package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Reqeust events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

var svcS3 *s3.S3
var svcDb *dynamodb.DynamoDB
var uploader *s3manager.Uploader
var bucketId string
var tableName string

func init() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		panic(err)
	}

	svcS3 = s3.New(sess)
	svcDb = dynamodb.New(sess)

	uploader = s3manager.NewUploader(sess)
	bucketId = os.Getenv("UPLOAD_IMAGE_STORAGE_ID")
	tableName = os.Getenv("IMAGE_TABLE")
}

func uploadS3(bucketId, fileExtension string, bodyReader io.Reader) (string, error) {
	fileUuid, errUuid := uuid.NewRandom()
	uploadedFileName := fmt.Sprintf("%s.%s", fileUuid.String(), fileExtension)

	if errUuid != nil {
		return "", fmt.Errorf("Unable to create random file name")
	}

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketId),
		Key:    aws.String(uploadedFileName),
		Body:   bodyReader,
	})

	if err != nil {
		// Print the error and exit.
		return "", fmt.Errorf("Unable to upload file to %q, %v", bucketId, err)
	}

	return uploadedFileName, nil
}

func createResponse(statusCode int, msg string) Response {
	return Response{
		StatusCode: statusCode,
		Body:       msg,
		Headers: map[string]string{
			"Content-Type": "plain/text",
		},
	}
}

type ImageItem struct {
	Id     string
	Status string
}

func createImageItem(id string) error {
	item := ImageItem{
		Id:     id,
		Status: "NEW",
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return fmt.Errorf("Got error marshalling new image item: ", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svcDb.PutItem(input)
	if err != nil {
		return fmt.Errorf("Got error calling PutItem: ", err)
	}

	return nil
}

func Handler(ctx context.Context, req Reqeust) (Response, error) {
	var extension string

	if strings.HasPrefix(req.Headers["Content-Type"], "image/") {
		extension = strings.TrimPrefix(req.Headers["Content-Type"], "image/")
	} else {
		return createResponse(500, "Wrong request content type"), nil
	}

	bodyStringReader := strings.NewReader(req.Body)
	reader := base64.NewDecoder(base64.StdEncoding, bodyStringReader)

	fileName, uploadErr := uploadS3(bucketId, extension, reader)
	createImageItem(fileName)

	if uploadErr != nil {
		return createResponse(500, uploadErr.Error()), nil
	}

	return createResponse(200, fileName), nil
}

func main() {
	lambda.Start(Handler)
}
