package clients

import (
	"errors"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbiface"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// DynamoDB mock
type dynamoDBMock struct {
	dynamodbiface.DynamoDBAPI
	putItemOutput *dynamodb.PutItemOutput
	getItemOutput *dynamodb.GetItemOutput
	err           error
}

func (v *dynamoDBMock) PutItemRequest(*dynamodb.PutItemInput) dynamodb.PutItemRequest {
	return dynamodb.PutItemRequest{
		Request: &aws.Request{
			Data:  v.putItemOutput,
			Error: v.err,
		},
	}
}

func (v *dynamoDBMock) GetItemRequest(*dynamodb.GetItemInput) dynamodb.GetItemRequest {
	return dynamodb.GetItemRequest{
		Request: &aws.Request{
			Data:  v.getItemOutput,
			Error: v.err,
		},
	}
}

// tests
func TestDynamoDBClient_AddURL(t *testing.T) {
	cases := []struct {
		name     string
		client   *dynamoDBMock
		shortURL string
		fullURL  string
		error    bool
	}{
		{
			name:   "Positive case",
			client: &dynamoDBMock{},
		},
		{
			name:   "Network error",
			client: &dynamoDBMock{err: errors.New("")},
			error:  true,
		},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			client := &DynamoDBClient{
				client: v.client,
			}
			err := client.AddURL(v.shortURL, v.fullURL)
			if v.error {
				assert.Error(t, err)
				log.Println(err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDynamoDBClient_GetURL(t *testing.T) {
	cases := []struct {
		name     string
		client   *dynamoDBMock
		shortURL string
		expected string
		error    bool
	}{
		{
			name: "Positive case",
			client: &dynamoDBMock{
				getItemOutput: &dynamodb.GetItemOutput{
					Item: map[string]dynamodb.AttributeValue{
						"fullUrl": {S: aws.String("http://example.com")},
					},
				},
			},
			shortURL: "abc",
			expected: "http://example.com",
		},
		{
			name: "Empty response",
			client: &dynamoDBMock{
				getItemOutput: &dynamodb.GetItemOutput{},
			},
			shortURL: "abc",
			expected: "http://example.com",
			error:    true,
		},
		{
			name:     "Network error",
			client:   &dynamoDBMock{err: errors.New("")},
			shortURL: "abc",
			expected: "http://example.com",
			error:    true,
		},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			client := &DynamoDBClient{
				client: v.client,
			}
			s, err := client.GetURL(v.shortURL)
			if v.error {
				assert.Error(t, err)
				log.Println(err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, v.expected, s)
			}
		})
	}
}
