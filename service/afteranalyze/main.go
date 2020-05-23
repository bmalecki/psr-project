package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/textract"
	"github.com/psr-project/uploadService/imageservice"
)

var textractSvc *textract.Textract
var imageTableSvc *imageservice.ImageTableService

func init() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	))

	textractSvc = textract.New(sess)
	imageTableSvc = imageservice.New(sess, os.Getenv("IMAGE_TABLE"))
}

func parseRecord(recordBody, nextToken *string) (*textract.GetDocumentTextDetectionInput, string, error) {
	var body TextractBody
	var message TextractMessage

	json.Unmarshal([]byte(*recordBody), &body)
	json.Unmarshal([]byte(*body.Message), &message)

	if *message.Status != "SUCCEEDED" {
		return nil, "", fmt.Errorf("Message status is not SUCCEEDED")
	}

	return &textract.GetDocumentTextDetectionInput{
		JobId:      message.JobId,
		MaxResults: aws.Int64(1000),
		NextToken:  nextToken,
	}, *message.DocumentLocation.S3ObjectName, nil
}

func Handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, record := range sqsEvent.Records {

		fmt.Printf("record: %s \n", record.Body)

		var fileName string
		var next *string

		for {
			input, documentLocation, errInput := parseRecord(&record.Body, next)
			if errInput != nil {
				return errInput
			}

			output, err := textractSvc.GetDocumentTextDetection(input)
			if err != nil {
				return err
			}

			next = output.NextToken
			fileName = documentLocation

			for _, block := range output.Blocks {
				if block.Text != nil && block.Confidence != nil {
					fmt.Printf("Text: %v | confidence: %v", *block.Text, *block.Confidence)
				}
			}

			if output.NextToken == nil {
				break
			}
		}

		if err := imageTableSvc.ReadyImageStatusItem(fileName); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	lambda.Start(Handler)
}
