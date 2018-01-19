package report

import (
	"fmt"

	"github.com/ingojaeckel/go-lambda-service-health/status"
)

type Reporter struct {
	S3Bucket    string
	S3KeyData   string
	S3KeyReport string
}

func (r Reporter) CreateReport(results []status.TimedResult, destination string) {
	var payload string

	for _, tr := range results {
		if tr.Success {
			payload += fmt.Sprintf("  UP: %10s %30s @ %4d ms, status: %3d\n", tr.Configuration.Name, tr.Configuration.URL, tr.TimeNanos, tr.StatusCode)
		} else {
			payload += fmt.Sprintf("DOWN: %10s %30s @ %4d ms, status: %3d\n", tr.Configuration.Name, tr.Configuration.URL, tr.TimeNanos, tr.StatusCode)
		}
	}

	fmt.Println("Report: ")
	fmt.Print(payload)

	// TODO write to s3
}

func (r Reporter) LoadExistingData() {
	// load data from S3
	from := fmt.Sprintf("s3//%s/%s", r.S3Bucket, r.S3KeyData)

	fmt.Printf("Loading from %s\n", from)
}
