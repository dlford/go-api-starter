package session

import (
	"api/cache"
	"api/models"
	"time"
)

func cacheAppkey(appkey *models.Appkey) {
	appkey.LastUsed = time.Now()
	cache.AddAppkey(appkey)
	models.DB.Save(appkey)
}
