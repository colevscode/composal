package track

import "time"

type Pattern struct {
	Begin    time.Duration `json:"begin"`
	Duration time.Duration `json:"duration"`
	Sample   string        `json:"sample"`
	Volume   float64       `json:"volume"`
	Pan      float64       `json:"pan"`
}

type Track struct {
	Type    string            `json:"type"`
	Samples map[string]string `json:"samples"`
	Pattern []Pattern         `json:"pattern"`
}
