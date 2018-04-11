package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var svc *dynamodb.DynamoDB

const tableName string = "emails"

type dynamoClient struct{}

var db = newDynamo()

func newDynamo() dynamoClient {
	return dynamoClient{}
}

func (db dynamoClient) setup() {
	db.createDynamoSession()
}

func (db dynamoClient) createDynamoSession() {
	key := os.Getenv("aws_access_key")
	secret := os.Getenv("aws_secret_key")
	region := os.Getenv("aws_region")

	awsConfig := &aws.Config{
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
		Region:      aws.String(region),
	}

	sess, err := session.NewSession(awsConfig)

	if err != nil {
		fmt.Println(err)
	}

	svc = dynamodb.New(sess)
}

func (db dynamoClient) isTableAvailable(tableName string) bool {

	tableInfo, err := db.getTable(tableName)
	if err != nil {
		return false
	}
	if *tableInfo.TableStatus == "ACTIVE" {
		return true
	}
	return false
}

func (db dynamoClient) getTable(tableName string) (*dynamodb.TableDescription, error) {

	params := &dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	}
	resp, err := svc.DescribeTable(params)

	if err != nil {
		return nil, err
	}
	table := *resp.Table
	return &table, nil
}

func (db dynamoClient) deleteEmail(key string, tableName string) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(key),
			},
		},
		TableName: aws.String(tableName),
	}

	_, err := svc.DeleteItem(input)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (db dynamoClient) insertEmail(email string, tableName string) error {
	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(email),
			},
		},
		TableName: aws.String(tableName),
	}
	_, err := svc.PutItem(input)

	if err != nil {
		return err
	}
	return nil
}

func (db dynamoClient) scanTable(tableName string) (interface{}, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}
	result, err := svc.Scan(input)

	if result != nil {
		var emails []emailAddr
		err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &emails)
		if err != nil {
			panic(err)
		}
		return emails, err
	}
	return nil, err
}
