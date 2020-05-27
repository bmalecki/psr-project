package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"github.com/psr-project/uploadService/imageservice"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Reqeust events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

var svcS3 *s3.S3
var imageTableService *imageservice.ImageTableService
var uploader *s3manager.Uploader
var bucketId string
var tableName string

func init() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	))

	svcS3 = s3.New(sess)
	uploader = s3manager.NewUploader(sess)
	bucketId = os.Getenv("UPLOAD_IMAGE_STORAGE_ID")

	imageTableService = imageservice.New(sess, os.Getenv("IMAGE_TABLE"))
}

func uploadS3(bucketId string, formData *FromData) (string, error) {
	fileUuid, errUuid := uuid.NewRandom()
	uploadedFileName := fmt.Sprintf("%s.%s", fileUuid.String(), formData.FileExtension)

	if errUuid != nil {
		return "", fmt.Errorf("Unable to create random file name")
	}

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucketId),
		Key:         aws.String(uploadedFileName),
		Body:        formData.FileReader,
		ContentType: aws.String(formData.ContentType),
	})

	if err != nil {
		return "", fmt.Errorf("Unable to upload file to %q, %v", bucketId, err)
	}

	return uploadedFileName, nil
}

type FromData struct {
	FileReader     io.Reader
	FileName       string
	ContentType    string
	FileExtension  string
	ForbiddenWords []string
}

func parseMultipartForm(contentType string, reader io.Reader) (*FromData, error) {
	_, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		return nil, err
	}
	mr := multipart.NewReader(reader, params["boundary"])

	var formData FromData

	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			return &formData, nil
		}
		if err != nil {
			return nil, err
		}
		slurp, err := ioutil.ReadAll(p)
		if err != nil {
			return nil, err
		}

		_, params, err := mime.ParseMediaType(p.Header.Get("Content-Disposition"))

		switch params["name"] {
		case "file":
			if strings.HasPrefix(p.Header.Get("Content-Type"), "image/") {
				formData.FileReader = bytes.NewReader(slurp)
				formData.FileName = params["filename"]
				formData.FileExtension = strings.TrimPrefix(p.Header.Get("Content-Type"), "image/")
				formData.ContentType = p.Header.Get("Content-Type")
			} else {
				return nil, fmt.Errorf("File is not an image")
			}
		case "forbiddenWords":
			formData.ForbiddenWords = append(formData.ForbiddenWords, string(slurp))
		}
	}
}

func createResponse(statusCode int, msg string) Response {
	return Response{
		StatusCode: statusCode,
		Body:       msg,
		Headers: map[string]string{
			"Content-Type":                     "plain/text",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}
}

func Handler(ctx context.Context, req Reqeust) (Response, error) {

	// See: https://github.com/aws/aws-lambda-go/issues/117
	contentType := req.Headers["Content-Type"]
	if len(contentType) == 0 {
		contentType = req.Headers["content-type"]
	}

	if !strings.HasPrefix(contentType, "multipart/form-data") {
		return createResponse(500, "Wrong request content type"), nil
	}

	bodyStringReader := strings.NewReader(req.Body)
	byteReader := base64.NewDecoder(base64.StdEncoding, bodyStringReader)
	formData, formErr := parseMultipartForm(contentType, byteReader)

	if formErr != nil {
		return createResponse(500, formErr.Error()), nil
	}

	objectId, uploadErr := uploadS3(bucketId, formData)

	if uploadErr != nil {
		return createResponse(500, uploadErr.Error()), nil
	}

	if err := imageTableService.CreateImageTableItem(objectId, formData.FileName, formData.ForbiddenWords); err != nil {
		return createResponse(500, err.Error()), nil
	}

	return createResponse(200, objectId), nil
}

func main() {
	lambda.Start(Handler)
}
