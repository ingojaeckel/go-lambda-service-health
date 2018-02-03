package report

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateReport(t *testing.T) {
	m1 := Measurement{"foo", 123, 200}
	m2 := Measurement{"bar", 456, 201}
	c1 := Check{9876543, []Measurement{m1, m2}}
	c2 := Check{9876544, []Measurement{m1, m2}}
	c3 := Check{9876545, []Measurement{m1, m2}}

	report := Report{Checks: []Check{c1, c2, c3}}

	fmt.Println(GenerateReport(report, time.Now().Add(-24 * time.Hour)))
}

func TestConvertTimes(t *testing.T) {
	ts := int64(1517366669)
	theTime := time.Unix(ts, 0)
	assert.Equal(t, "2018-01-30T18:44:29-08:00", theTime.Format(time.RFC3339))
}
