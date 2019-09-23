package model

import (
	"time"
)

type Message struct {
	ID        string    `json:"_id,omitempty"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
	User      User      `json:"user"`
}
