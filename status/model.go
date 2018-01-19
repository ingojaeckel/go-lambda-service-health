package status

type ServiceConfiguration struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type TimedResult struct {
	Configuration ServiceConfiguration
	Success       bool
	StatusCode    int
	TimeNanos     int64
}
