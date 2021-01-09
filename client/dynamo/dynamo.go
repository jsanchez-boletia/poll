package dynamo

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type dynamoClient struct {
	dynamodbiface.DynamoDBAPI
	table string
}

// New returns new client
func New(dynamo *dynamodb.DynamoDB, table string) dynamoClient {
	return dynamoClient{
		dynamo,
		table,
	}
}
