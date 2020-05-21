package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/textract"
	"github.com/psr-project/uploadService/imageservice"
)

var imageTableSvc *imageservice.ImageTableService
var textractSvc *textract.Textract
var analyzeImageQUrl string
var bucketId string
var textractSnsTopicArn string
var textractRoleArn string

func init() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	))

	imageTableSvc = imageservice.New(sess, os.Getenv("IMAGE_TABLE"))
	textractSvc = textract.New(sess)

	bucketId = os.Getenv("UPLOAD_IMAGE_STORAGE_ID")
	textractRoleArn = os.Getenv("TEXTRACT_ROLE_ARN")
	textractSnsTopicArn = os.Getenv("TEXTRACT_SNS_TOPIC_ARN")
}

func createDocumentTextDetectionInput(fileName string) *textract.StartDocumentTextDetectionInput {
	return &textract.StartDocumentTextDetectionInput{
		DocumentLocation: &textract.DocumentLocation{
			S3Object: &textract.S3Object{
				Bucket: aws.String(bucketId),
				Name:   aws.String(fileName),
			},
		},
		NotificationChannel: &textract.NotificationChannel{
			RoleArn:     aws.String(textractRoleArn),
			SNSTopicArn: aws.String(textractSnsTopicArn),
		},
	}
}

func Handler(ctx context.Context, sqsEvent events.SQSEvent) error {

	for _, record := range sqsEvent.Records {
		fileName := record.Body

		if err := imageTableSvc.ProcessingImageStatusItem(fileName); err != nil {
			return err
		}

		input := createDocumentTextDetectionInput(fileName)
		if _, err := textractSvc.StartDocumentTextDetection(input); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lambda.Start(Handler)
}
