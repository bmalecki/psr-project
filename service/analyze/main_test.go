package main

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/textract"
)

func TestTextRecognition(t *testing.T) {
	image := "542d908b-a0c6-4874-a6f6-94d3d346dcde.png"
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
			RoleArn:     aws.String("arn:aws:iam::429668857040:role/uploadservice-dev-AmazonTextractRole-EI3CQ46MSP47"),
			SNSTopicArn: aws.String("arn:aws:sns:us-east-1:429668857040:AmazonTextractSNSTopic"),
		},
	}

	if _, err := textractSvc.StartDocumentTextDetection(input); err != nil {
		t.Error(err)
	}

}
