package imageservice

import (
	"encoding/json"
	"fmt"
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
	imageTableService := New(sess, "uploadservice-dev-ImageTable-P9ENBK0K06ZA")

	// imageId := "f706094e-4681-4129-9350-74454c26a03b.png"

	// forbiddenWords, err := imageTableService.GetForbiddenWords(imageId)
	// if err != nil {
	// 	t.Errorf("%v", err)
	// }

	// fmt.Printf("%v\n", forbiddenWords)

	// var ofws []string
	// ofws = append(ofws, "aad")
	// ofws = append(ofws, "sada")
	// ofws = append(ofws, "fdsaf")

	// imageTableService.AddOccurredForbiddenWordsToItem(imageId, ofws)

	imageItemList, err := imageTableService.GetAllImageItems()

	if err != nil {
		fmt.Println(err.Error())
	}
	imageItemJson, err := json.Marshal(imageItemList)

	fmt.Println(string(imageItemJson))
}
