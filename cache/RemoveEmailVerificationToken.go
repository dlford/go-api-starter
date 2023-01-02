package cache

import "api/models"

func RemoveEmailVerificationToken(user *models.User) error {
	cacheKey := MakeEmailVerificationTokenKey(user.ID.String(), user.Email)
	redisErr := DB.Del(Ctx, cacheKey)
	if redisErr.Err() != nil {
		return redisErr.Err()
	}

	return nil
}
