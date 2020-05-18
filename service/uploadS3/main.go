package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/google/uuid"

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

func errResponse(msg string) Response {
	return Response{
		StatusCode: 500,
		Body:       msg,
		Headers: map[string]string{
			"Content-Type": "plain/text",
		},
	}
}

func Handler(ctx context.Context, req Reqeust) (Response, error) {
	bucketId := os.Getenv("BUCKET_ID")
	var extension string

	if strings.HasPrefix(req.Headers["Content-Type"], "image/") {
		extension = strings.TrimPrefix(req.Headers["Content-Type"], "image/")
	} else {
		return errResponse("Wrong request content type"), nil
	}

	bodyStringReader := strings.NewReader(req.Body)
	reader := base64.NewDecoder(base64.StdEncoding, bodyStringReader)

	_, uploadErr := uploadS3(bucketId, extension, reader)

	if uploadErr != nil {
		return errResponse(uploadErr.Error()), nil
	}

	resp := Response{
		StatusCode: 200,
		Body:       "OK",
		Headers: map[string]string{
			"Content-Type": "plain/text",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
