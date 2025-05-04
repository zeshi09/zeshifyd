package model

import (
	"time"
)

type Notification struct {
	Appname string    `json:"app_name"`
	Summary string    `json:"summary"`
	Body    string    `json:"body"`
	Icon    string    `json:"icon,omitepty"`
	Time    time.Time `json:"time"`
	Timeout int       `json:"timeout"` // milliseconds, 0 = notification doesn't disappear
}
