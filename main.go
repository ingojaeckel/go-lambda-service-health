package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ingojaeckel/go-lambda-service-health/config"
	"github.com/ingojaeckel/go-lambda-service-health/status"
)

// Handler Main handler function called by AWS Lambda.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	conf, err := config.LoadConfiguration("config.yaml")
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Failed to initialize: %s", err.Error()),
			StatusCode: 500,
		}, nil
	}

	resultChannel := make(chan status.TimedResult, 10)
	status.CheckResponseTime(conf, status.ServiceConfiguration{"foo", "https://google.com"}, resultChannel)
	defer close(resultChannel)

	tr := <-resultChannel
	if tr.Success {
		log.Printf("  UP: %10s %30s @ %4d ms, status: %3d\n", tr.Configuration.Name, tr.Configuration.URL, tr.TimeNanos/1000/1000, tr.StatusCode)
	} else {
		log.Printf("DOWN: %10s %30s @ %4d ms, status: %3d\n", tr.Configuration.Name, tr.Configuration.URL, tr.TimeNanos/1000/1000, tr.StatusCode)
	}

	return events.APIGatewayProxyResponse{
		Body:       "OK",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
