package session

import (
	"api/models"
	"errors"

	"github.com/alexedwards/argon2id"
)

func ValidatePasswordOrOtpCode(user models.User, password string, otpCode string) error {
	if user.OtpEnabled {
		return ValidateOtpCode(user, otpCode)
	} else {
		if password == "" {
			return errors.New("Password required")
		}

		var p models.Password
		models.DB.Model(&user).Association("Password").Find(&p)
		match, err := argon2id.ComparePasswordAndHash(password, p.Hash)
		if err != nil || !match {
			return errors.New("Invalid password")
		}
	}

	return nil
}
