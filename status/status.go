package status

import (
	"log"
	"net/http"
	"time"

	"github.com/ingojaeckel/go-lambda-service-health/config"
)

type TimedResult struct {
	Configuration config.ServiceConfiguration
	Success       bool
	StatusCode    int
	TimeNanos     int64
}


func CheckResponseTimes(configuration *config.Configuration) chan TimedResult {
	resultChannel := make(chan TimedResult, 10)

	for _, s := range configuration.Services {
		CheckResponseTime(configuration, s, resultChannel)
	}

	return resultChannel
}

func CheckResponseTime(configuration *config.Configuration, c config.ServiceConfiguration, out chan<- TimedResult) {
	log.Printf("checking response time for service %s @ %s\n", c.Name, c.URL)

	timeout := time.Duration(time.Duration(configuration.Timeout) * time.Second)
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
	if err == nil && resp != nil {
		return resp.StatusCode
	}
	return -1
}
