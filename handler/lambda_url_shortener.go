package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/artemnikitin/delex-code-example/clients"
	"github.com/artemnikitin/delex-code-example/model"

	"github.com/aws/aws-lambda-go/events"
	"github.com/satori/go.uuid"
)

// URLShortenerHandler represents Lambda entry point for creation of short URL
func URLShortenerHandler(request *model.URLShortenerRequest) (*model.URLShortenerResponse, error) {
	dynamo, err := clients.GetDynamoDBClient()
	if err != nil {
		return nil, err
	}

	short := uuid.NewV4().String()
	err = dynamo.AddURL(short, request.URL)
	if err != nil {
		return nil, err
	}

	return &model.URLShortenerResponse{ShortURL: short, Result: "success"}, nil
}

// URLRedirectHandler represents Lambda entry point for redirection to full URL by short URL
func URLRedirectHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	dynamo, err := clients.GetDynamoDBClient()
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}
	v, ok := request.PathParameters["short"]
	if !ok {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusNotImplemented}, err
	}

	URL, err := dynamo.GetURL(strings.Replace(v, "/", "", -1))
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadGateway}, err
	}

	log.Println(fmt.Sprintf("For short URL: %s, full URL: %s", v, URL))

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusMovedPermanently,
		Headers: map[string]string{
			"Location": URL,
		},
	}, nil
}
