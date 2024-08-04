package models

import "time"

type Subscription struct {
	ID        uint      `json:"-"`
	UserID    uint      `json:"user_id"`
	Topic     string    `json:"topic"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
   }