package clients

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbiface"
)

const (
	urlShortenerTable = "example-url-shortener-table"
)

// DynamoDBClient used to encapsulate DynamoDB specifics
type DynamoDBClient struct {
	client dynamodbiface.DynamoDBAPI
}

// GetDynamoDBClient is a constructor for DynamoDBClient
func GetDynamoDBClient() (*DynamoDBClient, error) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	cfg.Region = endpoints.EuWest1RegionID
	return &DynamoDBClient{client: dynamodb.New(cfg)}, nil
}

// AddURL adds full and generated short URL to DynamoDB
func (v *DynamoDBClient) AddURL(short, full string) error {
	req := v.client.PutItemRequest(&dynamodb.PutItemInput{
		Item: map[string]dynamodb.AttributeValue{
			"shortUrl": {S: aws.String(short)},
			"fullUrl":  {S: aws.String(full)},
			"ttl":      {N: aws.String(strconv.Itoa(getTTL()))},
		},
		TableName: aws.String(urlShortenerTable),
	})
	_, err := req.Send()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// GetURL returns full URL by short URL
func (v *DynamoDBClient) GetURL(short string) (string, error) {
	req := v.client.GetItemRequest(&dynamodb.GetItemInput{
		Key: map[string]dynamodb.AttributeValue{"shortUrl": {
			S: aws.String(short),
		}},
		TableName: aws.String(urlShortenerTable),
	})
	resp, err := req.Send()
	if err != nil {
		log.Println(err)
		return "", err
	}
	value, ok := resp.Item["fullUrl"]
	if !ok {
		log.Println(resp.Item)
		return "", errors.New("item doesn't have an attribute 'fullUrl'")
	}
	return *value.S, nil
}

func getTTL() int {
	return int(time.Now().Add(time.Hour * 24 * 20).Unix())
}
