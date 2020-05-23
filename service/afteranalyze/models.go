package main

type TextractDocumentLocation struct {
	S3ObjectName *string
	S3Bucket     *string
}

type TextractMessage struct {
	JobId            *string
	Status           *string
	API              *string
	Timestamp        *string
	DocumentLocation *TextractDocumentLocation
}

type TextractBody struct {
	Type             *string
	MessageId        *string
	TopicArn         *string
	Message          *string
	Timestamp        *string
	SignatureVersion *string
	Signature        *string
	SigningCertURL   *string
	UnsubscribeURL   *string
}
