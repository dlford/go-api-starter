package cache

import (
	"api/models"
	"encoding/json"
)

func GetAppkey(prefix string) (models.Appkey, error) {
	var appkey models.Appkey
	var cached models.AppkeyCached

	cacheKey := MakeAppkeyKey(prefix)

	appkeyJson, err := DB.Get(Ctx, cacheKey).Result()
	if err != nil || appkeyJson == "" {
		return appkey, err
	}

	err = json.Unmarshal([]byte(appkeyJson), &cached)
	if err != nil {
		return appkey, err
	}

	appkey = models.Appkey{
		ID:          cached.ID,
		UserId:      cached.UserId,
		User:        cached.User,
		Hash:        cached.Hash,
		Name:        cached.Name,
		Prefix:      cached.Prefix,
		Permissions: cached.Permissions,
		CreatedAt:   cached.CreatedAt,
		LastUsed:    cached.LastUsed,
	}

	return appkey, nil
}
