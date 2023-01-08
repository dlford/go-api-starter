package models

import (
	"encoding/json"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID      `json:"id" gorm:"primary_key;type:uuid;default:gen_random_uuid()" example:"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"`
	Email         string         `json:"email" gorm:"unique;not null;uniqueIndex" example:"john@yourdomain.com"`
	FirstName     string         `json:"first_name" gorm:"not null;" example:"John"`
	LastName      string         `json:"last_name" gorm:"not null;" example:"Doe"`
	OtpEnabled    bool           `json:"otp_enabled" gorm:"default:false;"`
	EmailVerified bool           `json:"email_verified" gorm:"default:false;"`
	Password      Password       `json:"-" gorm:"constraint:OnDelete:CASCADE;"`
	Otp           Otp            `json:"-" gorm:"constraint:OnDelete:CASCADE;"`
	RecoveryCodes []RecoveryCode `json:"-" gorm:"constraint:OnDelete:CASCADE;"`
	Sessions      []Session      `json:"-" gorm:"constraint:OnDelete:CASCADE;"`
	Appkeys       []Appkey       `json:"-" gorm:"constraint:OnDelete:CASCADE;"`
}

func (u User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}

type CreateUserInput struct {
	Email     string `json:"email" binding:"required,email" example:"john@yourdomain.com"`
	FirstName string `json:"first_name" binding:"required" example:"John"`
	LastName  string `json:"last_name" binding:"required" example:"Doe"`
	Password  string `json:"password" binding:"required,gte=6" example:"password1234"`
}

type SignInInput struct {
	Email    string `json:"email" binding:"required,email" example:"john@yourdomain.com"`
	Password string `json:"password" binding:"required" example:"password1234"`
	OtpCode  string `json:"otp_code" example:"123456"`
}

type UserEmailInput struct {
	Email string `json:"email" binding:"required,email" example:"john@yourdomain.com"`
}

type UserPasswordInput struct {
	Password string `json:"password" binding:"required" example:"password1234"`
}

type UpdateUserInput struct {
	Email     string `json:"email" binding:"required,email" example:"john@yourdomain.com"`
	FirstName string `json:"first_name" binding:"required" example:"John"`
	LastName  string `json:"last_name" binding:"required" example:"Doe"`
}

type UpdateUserPasswordInput struct {
	Password    string `json:"password" binding:"required" example:"password1234"`
	NewPassword string `json:"new_password" binding:"required,gte=6" example:"password12345"`
	OtpCode     string `json:"otp_code" example:"123456"`
}

type ResetUserPasswordInput struct {
	ID          string `json:"id" binding:"required" example:"b8f9c1c0-5b5e-4b4c-9c1c-0b5b5e4b4c9c"`
	NewPassword string `json:"new_password" binding:"required,gte=6" example:"password12345"`
	Token       string `json:"token" binding:"required" example:"token_from_reset_email"`
	OtpCode     string `json:"otp_code" example:"123456"`
}

type VerifyEmailAddressInput struct {
	ID          string `uri:"id" binding:"required,uuid" example:"b8f9c1c0-5b5e-4b4c-9c1c-0b5b5e4b4c9c"` // User ID
	Token       string `uri:"token" binding:"required"` // Token from email
	RedirectUrl string `uri:"redirect_url" validate:"optional"` // Redirect to this url after verification
}

type DeleteUserInput struct {
	OtpCode  string `json:"otp_code" example:"123456"`
	Password string `json:"password" example:"password1234"`
}
