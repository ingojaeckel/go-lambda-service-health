package status

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ingojaeckel/go-lambda-service-health/config"
)

func CheckResponseTime(c ServiceConfiguration, out chan<- TimedResult) {
	fmt.Printf("checking response time for service %s @ %s\n", c.Name, c.URL)

	timeout := time.Duration(config.TimeoutSeconds * time.Second)
	client := http.Client{Timeout: timeout}
	before := time.Now().UnixNano()
	resp, err := client.Get(c.URL)
	timeNanos := time.Now().UnixNano() - before

	out <- TimedResult{
		Configuration: c,
		Success:       isSuccess(resp, err),
		TimeNanos:     timeNanos,
		StatusCode:    getStatusCode(resp, err),
	}
}

func isSuccess(resp *http.Response, err error) bool {
	return err == nil && resp.StatusCode < 399 && resp.StatusCode >= 200
}

func getStatusCode(resp *http.Response, err error) int {
	if err != nil && resp != nil {
		return resp.StatusCode
	}
	return -1
}
