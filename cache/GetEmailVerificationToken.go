package cache

import (
	"api/models"
)

func GetEmailVerificationToken(user *models.User) (string, error) {
	query := MakeEmailVerificationTokenKey(user.ID.String(), user.Email)
	code, redisErr := DB.Get(Ctx, query).Result()
	if redisErr != nil {
		return "", redisErr
	}

	return code, nil
}
