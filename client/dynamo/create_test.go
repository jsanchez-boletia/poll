package dynamo

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/jsanchez-boletia/poll"
	"github.com/stretchr/testify/assert"
)

type dynamoCreateMock struct {
	dynamodbiface.DynamoDBAPI
	output      *dynamodb.PutItemOutput
	outputError error
}

func (d dynamoCreateMock) PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return d.output, d.outputError
}

func TestCreate(t *testing.T) {
	testCases := []struct {
		testName      string
		name          string
		event         string
		options       map[string]string
		expectedPoll  *poll.Poll
		expectedError error
		dynamoOutput  *dynamodb.PutItemOutput
		dynamoError   error
	}{
		{
			testName: "RegularCase",
			name:     "poll-01",
			event:    "event-01",
			options: map[string]string{
				"ANSWER-01": "GO",
				"ANSWER-02": "Perl",
				"ANSWER-03": "C",
				"ANSWER-04": "Rust",
			},
			expectedPoll: &poll.Poll{
				ID:             "1e75ccc5-9c01-4035-b65b-802fa6274b29",
				Name:           "poll-01",
				Action:         "",
				Active:         false,
				EventSubdomain: "event-01",
				Secuence:       "1",
				Answers: []poll.Answer{
					{
						ID:          "ANSWER-01",
						OptionLabel: "GO",
						Total:       0,
					},
					{
						ID:          "ANSWER-02",
						OptionLabel: "Perl",
						Total:       0,
					},
					{
						ID:          "ANSWER-03",
						OptionLabel: "C",
						Total:       0,
					},
					{
						ID:          "ANSWER-04",
						OptionLabel: "Rust",
						Total:       0,
					},
				},
			},
			dynamoOutput: &dynamodb.PutItemOutput{
				Attributes: map[string]*dynamodb.AttributeValue{
					"id": {
						S: aws.String("1e75ccc5-9c01-4035-b65b-802fa6274b29"),
					},
					"secuence": {
						N: aws.String("1"),
					},
					"name": {
						S: aws.String("poll-01"),
					},
					"action": {
						S: aws.String(""),
					},
					"active": {
						BOOL: aws.Bool(false),
					},
					"event_subdomain": {
						S: aws.String("event-subdomain"),
					},
					"answers": {
						L: []*dynamodb.AttributeValue{
							{
								M: map[string]*dynamodb.AttributeValue{
									"id": {
										S: aws.String("ANSWER-ID-0"),
									},
									"option_label": {
										S: aws.String("Ciudad de México"),
									},
									"total": {
										N: aws.String("100"),
									},
								},
							},
						},
					},
				},
			},
			dynamoError: nil,
		},
	}

	for _, c := range testCases {
		dyMock := dynamoCreateMock{
			output:      c.dynamoOutput,
			outputError: c.dynamoError,
		}

		client := dynamoClient{
			dyMock,
			"table",
		}
		t.Run(c.testName, func(t *testing.T) {
			poll, err := client.Create(c.name, c.event, c.options)

			if err != nil {
				assert.Nil(t, c.expectedPoll, poll)
			} else {
				assert.Equal(t, c.expectedPoll.Name, poll.Name)
				assert.Equal(t, c.expectedPoll.Action, poll.Action)
				assert.Equal(t, c.expectedPoll.Active, poll.Active)
				assert.Equal(t, c.expectedPoll.EventSubdomain, poll.EventSubdomain)
				assert.ElementsMatch(t, c.expectedPoll.Answers, poll.Answers)
				assert.Equal(t, c.expectedError, err)
			}
		})
	}
}
