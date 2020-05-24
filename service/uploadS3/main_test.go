package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

func TestUpload(t *testing.T) {
	resp, err := Handler(nil, Reqeust{})

	if resp.StatusCode != 200 {
		t.Errorf("Status code did not equal 200")
	}

	if err != nil {
		t.Errorf("Error occured")
	}

	var body map[string]interface{}

	json.Unmarshal([]byte(resp.Body), &body)

	if body["message"] != "Okay" {
		t.Errorf("Wrong message, was %v", body["message"])
	}
}

func createTestBucket(bucketIds ...string) string {
	var bucketId string

	if len(bucketIds) == 0 {
		bucketId = fmt.Sprintf("unittest-%s", uuid.New().String())
	} else {
		bucketId = bucketIds[0]
	}

	input := &s3.CreateBucketInput{
		Bucket: aws.String(bucketId),
	}

	if _, err := svcS3.CreateBucket(input); err != nil {
		panic(err)
	}

	return bucketId
}

func deleteBucket(bucketId string) {
	input := &s3.DeleteBucketInput{
		Bucket: aws.String(bucketId),
	}

	iter := s3manager.NewDeleteListIterator(svcS3, &s3.ListObjectsInput{
		Bucket: aws.String(bucketId),
	})

	if err := s3manager.NewBatchDeleteWithClient(svcS3).Delete(aws.BackgroundContext(), iter); err != nil {
		panic("Unable to delete objects from bucket")
	}

	if _, err := svcS3.DeleteBucket(input); err != nil {
		panic(err)
	}
}

// func TestS3Upload(t *testing.T) {
// 	bucketId := createTestBucket()
// 	defer deleteBucket(bucketId)

// 	content := "my request"
// 	name, err := uploadS3(bucketId, "txt", strings.NewReader(content))

// 	if err != nil {
// 		t.Error("Error during upload file to S3", err)
// 	}

// 	if !strings.HasSuffix(name, ".txt") {
// 		t.Errorf("Expected .txt suffix but was: %s", name)
// 	}
// }

func TestMulitpart(t *testing.T) {
	msg := `------WebKitFormBoundaryzCSuTbxRYc6IeBzz
Content-Disposition: form-data; name="file"; filename="test.json"
Content-Type: image/png

{
	"Type" : "Notification",
	"MessageId" : "303f31b7-8ec7-576b-b8cf-66f756632bd3",
}
------WebKitFormBoundaryzCSuTbxRYc6IeBzz
Content-Disposition: form-data; name="words"

aaaaa
------WebKitFormBoundaryzCSuTbxRYc6IeBzz
Content-Disposition: form-data; name="AAAA"

wdwe
------WebKitFormBoundaryzCSuTbxRYc6IeBzz--`

	formData, err := parseMultipartForm("multipart/form-data; boundary=----WebKitFormBoundaryzCSuTbxRYc6IeBzz", strings.NewReader(msg))

	if err != nil {
		log.Fatalln(err.Error())
	}

	bytes, _ := ioutil.ReadAll(formData.FileReader)
	s := string(bytes)

	fmt.Printf("%v", s)
}
