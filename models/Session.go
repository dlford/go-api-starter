package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID        uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:gen_random_uuid()" example:"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"`
	UserId    uuid.UUID `json:"-"`
	Salt      string    `json:"-" gorm:"not null;"`
	UserAgent string    `json:"user_agent" gorm:"not null;" example:"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36"`
	IpAddress string    `json:"ip_address" gorm:"not null; default: 'Unknown'" example:"127.0.0.1"`
	CreatedAt time.Time `json:"created_at" example:"2021-06-01T00:00:00Z"`
	ExpiresAt time.Time `json:"expires_at" example:"2021-06-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2021-06-01T00:00:00Z"`
}

type BearerTokenResponse struct {
	BearerToken string `json:"bearer_token" example:"Bearer ..."`
}