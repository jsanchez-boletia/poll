package dynamo

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	guuid "github.com/google/uuid"
	"github.com/jsanchez-boletia/poll"
)

func (d dynamoClient) Create(name, event string, options map[string]string) (*poll.Poll, error) {
	var newItem map[string]*dynamodb.AttributeValue
	var err error

	newPoll := poll.Poll{
		ID:             guuid.New().String(),
		Name:           name,
		EventSubdomain: event,
		Secuence:       fmt.Sprintf("%d", time.Now().Unix()),
	}

	if len(options) > 0 {
		newPoll.Answers = make([]poll.Answer, len(options))

		index := 0
		for opID, optLabel := range options {
			newPoll.Answers[index].ID = opID
			newPoll.Answers[index].OptionLabel = optLabel
			newPoll.Answers[index].Total = 0
			index++
		}
	}

	if newItem, err = dynamodbattribute.MarshalMap(newPoll); err != nil {
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      newItem,
		TableName: aws.String(d.table),
	}

	_, err = d.PutItem(input)
	if err != nil {
		return nil, err
	}

	return &newPoll, nil
}
