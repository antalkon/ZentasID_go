package logs

import "time"

type Log struct {
	ID        int       `json:"id"`
	IP        string    `json:"ip"`
	Method    string    `json:"method"`
	Path      string    `json:"path"`
	Timestamp time.Time `json:"timestamp"`
}
