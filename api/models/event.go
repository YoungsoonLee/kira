package models

import (
	"time"
)

// Event model
// Text is a event information.
// StartAt is event start date
// EndAt is event end date
type Event struct {
	Text    string    `json:"text" bson:"text"`
	StartAt time.Time `json:"start_at,omitempty" bson:"start_at"`
	EndAt   time.Time `json:"end_at,omitempty" bson:"end_at"`
}
