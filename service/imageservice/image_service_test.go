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
	imageTableService := New(sess, "uploadservice-dev-ImageTable-P9ENBK0K06ZA")

	imageId := "265fdcf7-da9f-4101-9d59-e11306c03906.png"

	err = imageTableService.ProcessingImageStatusItem(imageId)
	if err != nil {
		t.Errorf("%v", err)
	}

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

	// imageItemList, err := imageTableService.GetAllImageItems()

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// imageItemJson, err := json.Marshal(imageItemList)

	// fmt.Println(string(imageItemJson))
}
