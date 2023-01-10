package cache

import (
	"api/constants"
	"api/models"

	"github.com/gofrs/uuid"
)

func AddAccessToken(user *models.User, sessionId uuid.UUID, accessToken string) error {
	cacheKey := MakeAccessTokenKey(user.ID.String(), sessionId.String(), accessToken)

	redisErr := DB.Set(Ctx, cacheKey, user, constants.ACCESS_TOKEN_LIFETIME)
	if redisErr.Err() != nil {
		return redisErr.Err()
	}

	return nil
}
