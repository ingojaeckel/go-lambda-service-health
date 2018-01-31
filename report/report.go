package report

import (
	"encoding/json"
	"fmt"
	"time"
)

func GenerateReport(r Report) string {
	tsBytes, _ := json.Marshal(convertToTimeSeries(r.Checks))
	return fmt.Sprintf(htmlTemplate, string(tsBytes))
}

func convertToTimeSeries(checks []Check) []TimeSeries {
	serviceCount := len(checks[0].Measurements)
	ts := make([]TimeSeries, serviceCount)

	for i := 0; i < serviceCount; i++ {
		ts[i] = TimeSeries{Type: "scatter"}
		for _, c := range checks {
			ts[i].Name = c.Measurements[i].ServiceName
			ts[i].X = append(ts[i].X, convertTimestamp(c.Timestamp))
			ts[i].Y = append(ts[i].Y, c.Measurements[i].ResponseTime)
		}
	}
	return ts
}

func convertTimestamp(epoch int64) string {
	theTime := time.Unix(epoch, 0)
	return theTime.Format(time.RFC3339)
}

type TimeSeries struct {
	X    []string `json:"x"`
	Y    []int    `json:"y"`
	Name string   `json:"name"`
	Type string   `json:"type"`
}

type PageContent struct {
	data TimeSeries
}

const htmlTemplate = `
<html>
<head>
  <script src="https://cdn.plot.ly/plotly-latest.min.js"></script>
  <title>Service Health</title>
</head>
<body>
<h1>Service Health</h1>
<div id="myDiv"></div>
  <script>
  var data = %s;
  Plotly.newPlot('myDiv', data);
</script>
</body>
</html>`
