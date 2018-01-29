package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ingojaeckel/go-lambda-service-health/config"
	"github.com/ingojaeckel/go-lambda-service-health/report"
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

	reporter := report.Reporter{*conf}
	if _, err := reporter.GetExistingData(); err != nil {
		log.Printf("Failed to load existing report: %s", err.Error())
	} else {
		log.Print("Successful loaded report")
	}

	var check report.Check
	resultChannel := status.CheckResponseTimes(conf)

	for tr := range resultChannel {
		check.Measurements = append(check.Measurements, report.Measurement{
			ServiceName:  tr.Configuration.Name,
			ResponseTime: int(tr.TimeNanos / 1000 / 1000),
			StatusCode:   tr.StatusCode,
		})
	}

	var prevReport report.Report

	reporter.UpdateMeasurements(&prevReport, check)

	return events.APIGatewayProxyResponse{
		Body:       "OK",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
