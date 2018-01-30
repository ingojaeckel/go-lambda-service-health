package report

import "github.com/ingojaeckel/go-lambda-service-health/config"

type Reporter struct {
	Conf config.Configuration
}

type Measurement struct {
	ServiceName  string
	ResponseTime int
	StatusCode   int
}

type Check struct {
	Timestamp    int64
	Measurements []Measurement
}

type Report struct {
	Checks []Check
}
