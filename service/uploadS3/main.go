package main

import (
	"context"
	"fmt"
	"io"

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

func Handler(ctx context.Context, req Reqeust) (Response, error) {
	// var buf bytes.Buffer

	// body, err := json.Marshal(map[string]interface{}{
	// 	"message": "Okay",
	// })
	// if err != nil {
	// 	return Response{StatusCode: 404}, err
	// }
	// json.HTMLEscape(&buf, body)

	// var body string

	// if req.IsBase64Encoded {
	// 	body = "is"
	// } else {
	// 	body = "not is"
	// }

	resp := Response{
		StatusCode:      200,
		Body:            req.Body,
		IsBase64Encoded: true,
		Headers: map[string]string{
			"Content-Type": "content/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
