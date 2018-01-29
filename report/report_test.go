package report

import (
	"testing"

	"github.com/ingojaeckel/go-lambda-service-health/config"
	"github.com/ingojaeckel/go-lambda-service-health/status"
)

func TestReportResults(t *testing.T) {
	results := []status.TimedResult{
		{
			Configuration: config.ServiceConfiguration{Name: "Foo", URL: "http://foo.com"},
			Success:       true,
			TimeNanos:     1000,
			StatusCode:    200,
		},
		{
			Configuration: config.ServiceConfiguration{Name: "Bar", URL: "http://bar.com"},
			Success:       true,
			TimeNanos:     2000,
			StatusCode:    201,
		},
		{
			Configuration: config.ServiceConfiguration{Name: "Baz", URL: "http://example.com"},
			Success:       false,
			TimeNanos:     2000,
			StatusCode:    404,
		},
	}
	r := Reporter{}
	r.CreateReport(results, "foo")
}
