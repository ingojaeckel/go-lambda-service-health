package report

import (
	"fmt"
	"strconv"
	"strings"
	"log"
)

func parse(reportStr string) (*Report, error) {
	lines := strings.Split(reportStr, "\n")
	checks := make([]Check, 0, len(lines))

	for _, lineStr := range lines {
		if len(strings.TrimSpace(lineStr)) == 0 {
			// skip empty line
			continue
		}
		c, err := parseLine(lineStr)
		if err != nil {
			return nil, err
		}
		checks = append(checks, c)
	}

	return &Report{checks}, nil
}

// 123456789|service1,123,200 service2,true,123,201 service3,123,500 service4,-1,-1    // everything up except service4
func parseLine(lineStr string) (Check, error) {
	columns := strings.Split(lineStr, "|")
	ts, _ := strconv.ParseInt(columns[0], 10, 32)

	return Check{ts, parseMeasurements(columns[1])}, nil
}

func parseMeasurements(mstr string) []Measurement {
	substrings := strings.Split(mstr, " ")
	measurements := make([]Measurement, len(substrings))
	for i, substr := range substrings {
		mparts := strings.Split(substr, ",")
		if len(mparts) != 3 {
			log.Printf("Failed to parse measurement: '%s'. Skipping.\n", substr)
			continue
		}
		measurements[i] = Measurement{
			ServiceName:  mparts[0],
			ResponseTime: parseInt(mparts[1]),
			StatusCode:   parseInt(mparts[2]),
		}
	}
	return measurements
}

func parseInt(str string) int {
	value, _ := strconv.ParseInt(str, 10, 32)
	return int(value)
}

func (c Report) String() string {
	return fmt.Sprintf(mergeChecks(c.Checks, "\n"))
}

func (c Check) String() string {
	return fmt.Sprintf("%d|%s", c.Timestamp, mergeMeasurements(c.Measurements, " "))
}

func (c Measurement) String() string {
	return fmt.Sprintf("%s,%d,%d", c.ServiceName, c.ResponseTime, c.StatusCode)
}

func mergeChecks(checks []Check, glue string) string {
	var s string
	for i, x := range checks {
		s += x.String()
		if i < len(checks)-1 {
			s += glue
		}
	}
	return s
}

func mergeMeasurements(measurements []Measurement, glue string) string {
	var s string
	for i, x := range measurements {
		s += x.String()
		if i < len(measurements)-1 {
			s += glue
		}
	}
	return s
}
