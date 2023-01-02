package session

import (
	"api/models"
	"errors"

	"github.com/pquerna/otp/totp"
)

func ValidateOtpCode(user models.User, otpCode string) error {
	if user.OtpEnabled {
		var otp models.Otp
		models.DB.Model(&user).Association("Otp").Find(&otp)

		if otpCode == "" {
			return errors.New("OTP code required")
		}

		valid := totp.Validate(otpCode, otp.Secret)
		if !valid {
			var recoveryCodes []models.RecoveryCode
			models.DB.Model(&user).Association("RecoveryCodes").Find(&recoveryCodes)

			for _, r := range recoveryCodes {
				if otpCode == r.Code {
					models.DB.Save(&r)
					valid = true
					models.DB.Delete(&r)
					break
				}
			}

			if !valid {
				return errors.New("Invalid OTP code")
			}
		}
	}

	return nil
}
