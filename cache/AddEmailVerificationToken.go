package cache

import (
	"api/models"
	"fmt"
	"time"
)

func AddEmailVerificationToken(user *models.User, code string) error {
	lifetime, err := time.ParseDuration(fmt.Sprintf("%dh", 24 * 7))
	if err != nil {
		return err
	}

	cacheKey := MakeEmailVerificationTokenKey(user.ID.String(), user.Email)
	redisErr := DB.Set(Ctx, cacheKey, code, lifetime)
	if redisErr.Err() != nil {
		return redisErr.Err()
	}

	return nil
}
