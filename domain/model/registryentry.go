package model

import "time"

type Registry struct {
	Transaction string    `json:"transaction"`
	Date        time.Time `json:"date"`
	Duration    float64   `json:"duration"`
}
