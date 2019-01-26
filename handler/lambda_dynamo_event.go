package handler

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

func InspectDynamoDBEventHandler(e events.DynamoDBEvent) {
	log.Println("Total number of changes:", len(e.Records))
	for _, record := range e.Records {
		log.Println(fmt.Printf("Processing event ID %s, type %s", record.EventID, record.EventName))
		// Logic for processing event goes here
	}
}
