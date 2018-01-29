package report

import (
	"fmt"
	"strconv"
	"strings"
)

type Measurement struct {
	serviceName  string
	responseTime int
	statusCode   int
}

type Check struct {
	timestamp    int
	measurements []Measurement
}

type Report struct {
	checks []Check
}

func (c Check) String() string {
	return fmt.Sprintf("%d|%s", c.timestamp, merge(c.measurements))
}

func merge(m []Measurement) string {
	var s string

	for i, x := range m {
		s += x.String()
		if i < len(m)-1 {
			s += " "
		}
	}
	return s
}

func parse(reportStr string) (*Report, error) {
	lines := strings.Split(reportStr, "\n")
	checks := make([]Check, len(lines))

	for i, lineStr := range lines {
		c, err := parseLine(lineStr)
		if err != nil {
			return nil, err
		}
		checks[i] = c
	}

	return &Report{checks}, nil
}

// 123456789|service1,123,200 service2,true,123,201 service3,123,500 service4,-1,-1    // everything up except service4
func parseLine(lineStr string) (Check, error) {
	columns := strings.Split(lineStr, "|")

	return Check{parseInt(columns[0]), parseMeasurements(columns[1])}, nil
}

func parseMeasurements(mstr string) []Measurement {
	substrings := strings.Split(mstr, " ")
	measurements := make([]Measurement, len(substrings))
	for i, substr := range substrings {
		mparts := strings.Split(substr, ",")
		measurements[i] = Measurement{
			serviceName:  mparts[0],
			responseTime: parseInt(mparts[1]),
			statusCode:   parseInt(mparts[2]),
		}
	}
	return measurements
}

func parseInt(str string) int {
	value, _ := strconv.ParseInt(str, 10, 32)
	return int(value)
}

func (c Measurement) String() string {
	return fmt.Sprintf("%s,%d,%d", c.serviceName, c.responseTime, c.statusCode)
}
