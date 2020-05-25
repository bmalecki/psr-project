package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

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
		var imageId string
		var next *string
		var forbiddenWords []string
		occurredForbiddenWordsMap := make(map[string]bool)

		for {
			input, documentLocation, errInput := parseRecord(&record.Body, next)
			if errInput != nil {
				return errInput
			}

			output, err := textractSvc.GetDocumentTextDetection(input)
			if err != nil {
				return err
			}

			if len(imageId) == 0 {
				imageId = documentLocation
				forbiddenWords, err = imageTableSvc.GetForbiddenWords(imageId)
			}
			if err != nil {
				return err
			}

			next = output.NextToken

			for _, block := range output.Blocks {
				if block.Text != nil && block.Confidence != nil && *block.Confidence > 75 {
					for _, forbiddenWord := range forbiddenWords {
						if strings.Contains(*block.Text, forbiddenWord) {
							occurredForbiddenWordsMap[forbiddenWord] = true
						}
					}
				}
			}

			if output.NextToken == nil {
				break
			}
		}

		var occurredForbiddenWords []string
		for k := range occurredForbiddenWordsMap {
			occurredForbiddenWords = append(occurredForbiddenWords, k)
		}

		if err := imageTableSvc.AddOccurredForbiddenWordsToItem(imageId, occurredForbiddenWords); err != nil {
			return err
		}

		if err := imageTableSvc.ReadyImageStatusItem(imageId); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	lambda.Start(Handler)
}
