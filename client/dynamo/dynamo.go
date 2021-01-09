package dynamo

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type dynamoClient struct {
	dynamodbiface.DynamoDBAPI
	table string
}

// New returns new client
func New(dynamo dynamodbiface.DynamoDBAPI, table string) dynamoClient {
	return dynamoClient{
		dynamo,
		table,
	}
}
