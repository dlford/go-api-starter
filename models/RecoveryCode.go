package models

import (
	"github.com/gofrs/uuid"
)

type RecoveryCode struct {
	ID     uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	UserId uuid.UUID
	Code   string `json:"code" gorm:"not null;"`
}
