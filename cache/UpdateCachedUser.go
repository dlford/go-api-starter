package cache

import (
	"api/models"
)

func UpdateCachedUser(user *models.User) {
	query := MakeAccessTokenKey(user.ID.String(), "*", "*")
	matches := DB.Keys(Ctx, query).Val()

	for _, match := range matches {
		ttl := DB.TTL(Ctx, match)
		DB.Set(Ctx, match, user, ttl.Val())
	}
}
