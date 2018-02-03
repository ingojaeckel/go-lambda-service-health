package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ingojaeckel/go-lambda-service-health/config"
	"github.com/ingojaeckel/go-lambda-service-health/report"
	"github.com/ingojaeckel/go-lambda-service-health/status"
	"time"
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
	existingReport, err := reporter.GetExistingData()
	if err != nil {
		// Continue anyway..
		log.Printf("Failed to load existing report: %s\n", err.Error())
		existingReport = &report.Report{}
	} else {
		log.Println("Successfully loaded report.")
	}

	check := report.Check{Timestamp: time.Now().Unix()}
	resultChannel := status.CheckResponseTimes(conf)

	for tr := range resultChannel {
		check.Measurements = append(check.Measurements, report.Measurement{
			ServiceName:  tr.Configuration.Name,
			ResponseTime: int(tr.TimeNanos / 1000 / 1000),
			StatusCode:   tr.StatusCode,
		})
	}

	log.Println("Uploading report..")
	if err := reporter.UpdateMeasurements(existingReport, check); err != nil {
		log.Printf("Failed to upload measurements: %s", err.Error())
	}

	return events.APIGatewayProxyResponse{
		Body:       "OK",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
