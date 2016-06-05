package main

import (
	// "fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var ddbClient = dynamodb.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})

// Put in AWS DynamoDB
func putToDynamoDB(mp *MeasurementProtocol) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("tblname"),
	}

	_, err := ddbClient.PutItem(input)

	if err != nil {
		// Send has failed - drop the record and proceed.
		log.Printf("Sender failed to put record: %s, %v, %#v\n", err, err, err)
	} else {
		log.Printf("Sending successful.\n")
	}

	return nil
}

func init() {
	// config = aws.DefaultConfig
	// add to putFunc registry
	putFuncRegistry["dynamodb"] = putToDynamoDB
}
