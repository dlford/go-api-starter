package cache

import (
	"api/models"
	"errors"
	"time"

	"github.com/google/uuid"
)

func GetUserFromAccessToken(bearerToken string) (models.User, int, error) {
	query := MakeAccessTokenKey("*", "*", bearerToken)
	matches := DB.Keys(Ctx, query).Val()

	if len(matches) > 0 {
		bearerCacheKey := matches[0]

		var user models.User
		err := DB.Get(Ctx, bearerCacheKey).Scan(&user)
		if err == nil || user.ID != uuid.Nil {
			ttlRaw := DB.TTL(Ctx, matches[0])
			if ttlRaw == nil || ttlRaw.Val() < 0 {
				return models.User{}, 0, errors.New("Unauthorized")
			}
			ttl := int(time.Duration(ttlRaw.Val().Seconds()))

			return user, ttl, nil
		}
	}

	return models.User{}, 0, errors.New("Unauthorized")
}
