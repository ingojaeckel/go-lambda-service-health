package merge

import "fmt"

type Measurement struct {
	serviceName  string
	responseTime int
	statusCode   int
}

type Check struct {
	timestamp    int
	measurements []Measurement
}

func (c Check) String() string {
	return fmt.Sprintf("%s %s", c.timestamp, merge(c.measurements))
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

func (c Measurement) String() string {
	return fmt.Sprintf("%s,%d,%d", c.serviceName, c.responseTime, c.statusCode)
}
