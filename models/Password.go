package models

import (
	"github.com/gofrs/uuid"
)

type Password struct {
	ID     uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	UserId uuid.UUID
	Hash   string `json:"hash" gorm:"not null;"`
}
