package imageservice

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func TestDynamoDb(t *testing.T) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		panic(err)
	}
	imageTableService := New(sess, "uploadservice-dev-ImageTable-1UOQKMF3IKZAI")

	if err := imageTableService.ProcessingImageStatusItem("cab9b42b-2141-4e20-89ca-fe000e33b483.png"); err != nil {
		t.Errorf("%v", err)
	}
}
