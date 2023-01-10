package models

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
)

type AppkeyPermissions struct {
	CanViewUserInfo bool `json:"can_view_user_info" gorm:"default:false;"`
	CanEditUserInfo bool `json:"can_edit_user_info" gorm:"default:false;"`
}

type Appkey struct {
	ID          uuid.UUID         `json:"id" gorm:"primary_key;type:uuid;default:gen_random_uuid()" example:"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"`
	UserId      uuid.UUID         `json:"-" gorm:"not null;"`
	User        User              `json:"-"`
	Hash        string            `json:"-" gorm:"not null;"`
	Name        string            `json:"name" gorm:"not null;" example:"My App"`
	Prefix      string            `json:"prefix" gorm:"unique;not null;" example:"1xauog0QhGi6MUzo"`
	Permissions AppkeyPermissions `json:"permissions" gorm:"embedded"`
	CreatedAt   time.Time         `json:"created_at" example:"2021-06-01T00:00:00Z"`
	LastUsed    time.Time         `json:"last_used" gorm:"autoCreateTime" example:"2021-06-01T00:00:00Z"`
}

type AppkeyCached struct {
	ID          uuid.UUID         `json:"id" gorm:"primary_key;type:uuid;default:gen_random_uuid()" example:"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"`
	UserId      uuid.UUID         `json:"user_id" gorm:"not null;"`
	User        User              `json:"user"`
	Hash        string            `json:"hash" gorm:"not null;"`
	Name        string            `json:"name" gorm:"not null;" example:"My App"`
	Prefix      string            `json:"prefix" gorm:"unique;not null;" example:"1xauog0QhGi6MUzo"`
	Permissions AppkeyPermissions `json:"permissions" gorm:"embedded"`
	CreatedAt   time.Time         `json:"created_at" example:"2021-06-01T00:00:00Z"`
	LastUsed    time.Time         `json:"last_used" example:"2021-06-01T00:00:00Z"`
}

func (u AppkeyCached) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u *AppkeyCached) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}

type AppkeyAuth struct {
	IsAppkey    bool
	Permissions AppkeyPermissions
}

type CreateAppkeyInput struct {
	Name        string            `json:"name" binding:"required" example:"My App"`
	Permissions AppkeyPermissions `json:"permissions" binding:"required"`
	OtpCode     string            `json:"otp_code" example:"123456"`
	Password    string            `json:"password" example:"password1234"`
}

type CreateAppkeyResponse struct {
	Key string `json:"key" example:"Bearer ..."`
}
