package main

import (
	"log"

	"github.com/artemnikitin/delex-code-example/handler"

	"github.com/aws/aws-lambda-go/lambda"
	_ "go.elastic.co/apm/module/apmlambda"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	lambda.Start(handler.URLShortenerHandler)
}
