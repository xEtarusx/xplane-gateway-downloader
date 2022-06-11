package types

import "time"

type Release struct {
	Version string    `json:"Version"`
	Date    time.Time `json:"Date"`
}
