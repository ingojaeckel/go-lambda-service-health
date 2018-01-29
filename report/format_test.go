package report

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReportSerializationMeasurement(t *testing.T) {
	m := Measurement{"foo", 123, 200}
	assert.Equal(t, "foo,123,200", m.String())
}

func TestReportSerializationCheck(t *testing.T) {
	m1 := Measurement{"foo", 123, 200}
	m2 := Measurement{"bar", 456, 201}
	c := Check{9876543, []Measurement{m1, m2}}

	assert.Equal(t, "9876543|foo,123,200 bar,456,201", c.String())
}

func TestReportSerialization(t *testing.T) {
	m1 := Measurement{"foo", 123, 200}
	m2 := Measurement{"bar", 456, 201}
	c1 := Check{9876543, []Measurement{m1, m2}}
	c2 := Check{9876544, []Measurement{m1, m2}}
	c3 := Check{9876545, []Measurement{m1, m2}}

	report := Report{checks: []Check{c1, c2, c3}}

	reportStr := `9876543|foo,123,200 bar,456,201
9876544|foo,123,200 bar,456,201
9876545|foo,123,200 bar,456,201`

	assert.Equal(t, reportStr, report.String())
}

func TestParseReport(t *testing.T) {
	reportStr := `9876543|foo,123,200 bar,456,201
9876544|foo,123,200 bar,456,201
9876545|foo,123,200 bar,456,201`

	r, err := parse(reportStr)
	assert.Nil(t, err)
	assert.NotNil(t, r)

	assert.Equal(t, 3, len(r.checks))
	assert.Equal(t, 2, len(r.checks[0].measurements))
	assert.Equal(t, "bar", r.checks[0].measurements[1].serviceName)
	assert.Equal(t, 456, r.checks[0].measurements[1].responseTime)
	assert.Equal(t, 201, r.checks[0].measurements[1].statusCode)
}