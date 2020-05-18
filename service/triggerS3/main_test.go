package main

import (
	"testing"
)

func TestPushSQS(t *testing.T) {
	qURL := "https://sqs.us-east-1.amazonaws.com/429668857040/AnalyzeImageQueue"

	pushRecord(qURL)
}
