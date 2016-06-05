package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

var kin = kinesis.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})

// Put in AWS DynamoDB
func putToKinesis(mp *MeasurementProtocol) error {
	params := &kinesis.PutRecordInput{
		Data:                      []byte("PAYLOAD"),          // Required
		PartitionKey:              aws.String("PartitionKey"), // Required
		StreamName:                aws.String("StreamName"),   // Required
		ExplicitHashKey:           aws.String("HashKey"),
		SequenceNumberForOrdering: aws.String("SequenceNumber"),
	}

	_, err := kin.PutRecord(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Get error details
			log.Println("Error:", awsErr.Code(), awsErr.Message())

			// Prints out full error message, including original error if there was one.
			log.Println("Error:", awsErr.Error())

			// Get original error
			// if origErr := awsErr.OrigErr(); origErr != nil {
			// 	// operate on original error.
			// }
		} else {
			fmt.Println(err.Error())
		}
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	return nil
}

func init() {
	// config = aws.DefaultConfig
	// add to putFunc registry
	putFuncRegistry["kinesis"] = putToKinesis
}
