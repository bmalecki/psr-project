package main

import (
	"fmt"
	"testing"
)

func TestTextractSqsEvent(t *testing.T) {

	recordBody := `{
		"Type" : "Notification",
		"MessageId" : "995966fe-90eb-5b48-80ce-31bcd67cd021",
		"TopicArn" : "arn:aws:sns:us-east-1:429668857040:AmazonTextractSNSTopic",
		"Message" : "{\"JobId\":\"8d34e8bdc77378a40812fc76097c67e35aba9a4169bae14ea691f744ea654686\",\"Status\":\"SUCCEEDED\",\"API\":\"StartDocumentTextDetection\",\"Timestamp\":1590182999890,\"DocumentLocation\":{\"S3ObjectName\":\"9a0e6077-d226-4de5-86e2-219c9f9213a4.png\",\"S3Bucket\":\"uploadservice-dev-uploadimagestorage-13vmpn57dtka4\"}}",
		"Timestamp" : "2020-05-22T21:29:59.951Z",
		"SignatureVersion" : "1",
		"Signature" : "aVbzc+AWwgQflu/dGwlRoUxQRaG2oAMa70hSWArLC5G7If1Cel2lnQlEFABBVRg+TBD35LwEHwoca7Ktl9JzktndXgO0YaGyKO/8ge8khh+gwAdulQxwx77ttHo0KuKg1U2a6BZh8I34LiWBV7lNtxi6xhZuSHw/chKOmCduqCXfsHXF7sBUuBZqlp3s9LCgC4LlGDIGF6+HV+UD9eUnyIDquMcrHMmPQ2qtHZdWPvQR5VcYBcE2iXjcfRhVTtKZO1orO0gGjf1a7UP7xdPaqQyqSRbpAotvBAPxjxaBne0V0e0nG+McWvyz6GeD07JYahl/bkj+RZ4otI/UgIA1sQ==",
		"SigningCertURL" : "https://sns.us-east-1.amazonaws.com/SimpleNotificationService-a86cb10b4e1f29c941702d737128f7b6.pem",
		"UnsubscribeURL" : "https://sns.us-east-1.amazonaws.com/?Action=Unsubscribe&SubscriptionArn=arn:aws:sns:us-east-1:429668857040:AmazonTextractSNSTopic:7ee040c0-9b1f-46af-8b64-b84495f333e0"
		}`

	_, a, _ := parseRecord(&recordBody, nil)

	fmt.Printf("Test: %v \n", a)

}
