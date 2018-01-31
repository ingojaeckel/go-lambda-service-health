package report

import "github.com/ingojaeckel/go-lambda-service-health/config"

type Reporter struct {
	Conf config.Configuration
}

type Measurement struct {
	ServiceName  string `json:"s"`
	ResponseTime int    `json:"t"`
	StatusCode   int    `json:"st"`
}

type Check struct {
	Timestamp    int64         `json:"ts"`
	Measurements []Measurement `json:"d"`
}

type Report struct {
	Checks []Check `json:"checks"`
}
