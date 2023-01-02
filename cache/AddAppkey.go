package cache

import (
	"api/constants"
	"api/models"
)

func AddAppkey(appkey *models.Appkey) error {
	cacheKey := MakeAppkeyKey(appkey.Prefix)

	cached := models.AppkeyCached{
		ID:          appkey.ID,
		UserId:      appkey.UserId,
		User:        appkey.User,
		Hash:        appkey.Hash,
		Name:        appkey.Name,
		Prefix:      appkey.Prefix,
		Permissions: appkey.Permissions,
		CreatedAt:   appkey.CreatedAt,
		LastUsed:    appkey.LastUsed,
	}

	redisErr := DB.Set(Ctx, cacheKey, cached, constants.APPKEY_CACHE_LIFETIME)
	if redisErr.Err() != nil {
		return redisErr.Err()
	}

	return nil
}
