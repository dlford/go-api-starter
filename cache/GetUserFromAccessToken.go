package cache

import (
	"api/models"
	"encoding/json"
	"errors"
	"time"
)

func GetUserFromAccessToken(bearerToken string) (models.User, int, error) {
	query := MakeAccessTokenKey("*", "*", bearerToken)
	matches := DB.Keys(Ctx, query).Val()

	if len(matches) > 0 {
		bearerCacheKey := matches[0]

		userJson := DB.Get(Ctx, bearerCacheKey).Val()
		if userJson != "" {

			var user models.User
			err := json.Unmarshal([]byte(userJson), &user)
			if err != nil {
				return models.User{}, 0, err
			}

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
