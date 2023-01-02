package cache

import "api/models"

func GetPasswordResetToken(user *models.User) (string, error) {
	query := MakePasswordResetTokenKey(user.ID.String())
	code, redisErr := DB.Get(Ctx, query).Result()
	if redisErr != nil {
		return "", redisErr
	}

	DB.Del(Ctx, user.ID.String()+":PasswordReset")

	return code, nil
}
