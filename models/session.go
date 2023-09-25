package models

import "time"

type Session struct {
	Token     string    `json:"token"`
	User      int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}
