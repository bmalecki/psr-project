package imageservice

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type imageItem struct {
	Id          string
	ImageStatus string
	Name        string
}

type ImageTableService struct {
	svcDb     *dynamodb.DynamoDB
	tableName string
}

func New(sess *session.Session, tableName string) *ImageTableService {
	return &ImageTableService{
		dynamodb.New(sess),
		tableName,
	}
}

func (it *ImageTableService) CreateImageTableItem(id, name string) error {
	item := imageItem{
		Id:          id,
		ImageStatus: "NEW",
		Name:        name,
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return fmt.Errorf("Got error marshalling new image item: %v", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(it.tableName),
	}

	_, err = it.svcDb.PutItem(input)
	if err != nil {
		return fmt.Errorf("Got error calling PutItem: %v", err)
	}

	return nil
}

func (it *ImageTableService) updateImageTableItem(id, column, value string) error {
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":s": {
				S: aws.String(value),
			},
		},
		TableName: aws.String(it.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(id),
			},
		},
		UpdateExpression: aws.String(fmt.Sprintf("set %s = :s", column)),
	}

	if _, err := it.svcDb.UpdateItem(input); err != nil {
		return err
	}
	return nil
}

func (it *ImageTableService) ProcessingImageStatusItem(id string) error {
	return it.updateImageTableItem(id, "ImageStatus", "PROCESSING")
}

func (it *ImageTableService) ReadyImageStatusItem(id string) error {
	return it.updateImageTableItem(id, "ImageStatus", "READY")
}
