package models

import "github.com/gofrs/uuid"

type Otp struct {
	ID     uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	UserId uuid.UUID
	Secret string `json:"secret" gorm:"not null;"`
}

type EnableOtpInput struct {
	OtpCode string `json:"otp_code" binding:"required" example:"123456"`
}

type SetupOtpResponse struct {
	Secret string `json:"secret" example:"JBSWY3DPEHPK3PXP"`
	QrCode string `json:"qr_code" example:"data:image/png;base64,..."`
}

type EnableOtpResponse struct {
	RecoveryCodes [10]string `json:"recovery_codes" example:"12345678,12345678,12345678,12345678,12345678"`
}
