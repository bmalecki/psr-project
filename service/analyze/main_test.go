package main

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/textract"
)

func TestTextRecognition(t *testing.T) {
	image := "e5978e6e-f7d9-4774-b4c3-805df99c47df.png"
	bucket := "uploadservice-dev-uploadimagestorage-13gnkse112ttl"
	fmt.Printf(image)

	// input := &textract.DetectDocumentTextInput{
	// 	Document: &textract.Document{
	// 		S3Object: &textract.S3Object{
	// 			Bucket: aws.String(bucket),
	// 			Name:   aws.String(image),
	// 		},
	// 	},
	// }

	// req, resp := textractSvc.DetectDocumentTextRequest(input)

	// err := req.Send()
	// if err == nil { // resp is now filled
	// 	fmt.Println(resp)
	// }

	// fmt.Printf(resp.String())

	input := &textract.StartDocumentTextDetectionInput{
		DocumentLocation: &textract.DocumentLocation{
			S3Object: &textract.S3Object{
				Bucket: aws.String(bucket),
				Name:   aws.String(image),
			},
		},
		NotificationChannel: &textract.NotificationChannel{
			RoleArn:     aws.String("arn:aws:iam::429668857040:role/uploadservice-dev-us-east-1-lambdaRole"),
			SNSTopicArn: aws.String("arn:aws:sns:us-east-1:429668857040:AmazonTextractSNSTopic"),
		},
	}

	if _, err := textractSvc.StartDocumentTextDetection(input); err != nil {
		t.Error(err)
	}

}
