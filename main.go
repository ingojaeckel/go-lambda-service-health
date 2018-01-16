package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler Main handler function called by AWS Lambda.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	return events.APIGatewayProxyResponse{
		Body:       "Hello World",
		StatusCode: 200,
	}, nil

}

func main() {
	lambda.Start(Handler)
}
