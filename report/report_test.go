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

	fmt.Println(GenerateReport(report, time.Now().Add(-24*time.Hour)))
}

func TestConvertTimes(t *testing.T) {
	ts := int64(1517366669)
	theTime := time.Unix(ts, 0)
	assert.Equal(t, "2018-01-30T18:44:29-08:00", theTime.Format(time.RFC3339))
}

func TestReportPreparation(t *testing.T) {
	m1 := Measurement{"foo", 123, 200}
	m2 := Measurement{"bar", 456, 201}
	c1 := Check{9876543, []Measurement{m1, m2}}
	c2 := Check{9876544, []Measurement{m1, m2}}
	c3 := Check{9876545, []Measurement{m1, m2}}

	report := Report{Checks: []Check{c1, c2, c3}}

	timestamps := convertToTimeSeries(report.Checks, time.Now())
	assert.Equal(t, 2, len(timestamps))
	assert.Equal(t, 0, len(timestamps[0].X))
	assert.Equal(t, 0, len(timestamps[1].X))
}

func TestReportPreparation2(t *testing.T) {
	now := time.Now()
	nowSeconds := now.Unix()
	nowSecondsStr := now.Format(time.RFC3339)

	m1 := Measurement{"foo", 123, 200}
	m2 := Measurement{"bar", 456, 201}

	c1 := Check{9876543, []Measurement{m1, m2}}
	c2 := Check{9876544, []Measurement{m1, m2}}
	c3 := Check{9876545, []Measurement{m1, m2}}
	c4 := Check{nowSeconds - 2, []Measurement{m1, m2}}
	c5 := Check{nowSeconds - 1, []Measurement{m1, m2}}
	c6 := Check{nowSeconds - 0, []Measurement{m1, m2}}

	report := Report{Checks: []Check{c1, c2, c3, c4, c5, c6}}

	timestamps := convertToTimeSeries(report.Checks, time.Now().Add(-24*time.Hour))
	assert.Equal(t, 2, len(timestamps))
	assert.Equal(t, 3, len(timestamps[0].X))
	assert.Equal(t, 3, len(timestamps[1].X))
	assert.Equal(t, nowSecondsStr, timestamps[0].X[2])
	assert.Equal(t, time.Unix(nowSeconds - 2, 0).Format(time.RFC3339), timestamps[0].X[0])
	assert.Equal(t, time.Unix(nowSeconds - 1, 0).Format(time.RFC3339), timestamps[0].X[1])
	assert.Equal(t, time.Unix(nowSeconds - 0, 0).Format(time.RFC3339), timestamps[0].X[2])
}
