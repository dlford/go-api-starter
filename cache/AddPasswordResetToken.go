package cache

import (
	"api/models"
	"time"
)

func AddPasswordResetToken(user *models.User, code string) error {
	lifetime, err := time.ParseDuration("15m")
	if err != nil {
		return err
	}

	cacheKey := MakePasswordResetTokenKey(user.ID.String())
	redisErr := DB.Set(Ctx, cacheKey, code, lifetime)
	if redisErr.Err() != nil {
		return redisErr.Err()
	}

	return nil
}
