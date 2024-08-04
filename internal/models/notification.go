package models

import "time"

type Notification struct {
	ID        uint      `json:"-"`
	UserID    uint      `json:"user_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
   }