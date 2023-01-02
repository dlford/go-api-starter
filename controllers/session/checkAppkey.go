package session

import (
	"api/cache"
	"api/models"
	"errors"
	"fmt"
	"strings"

	"github.com/alexedwards/argon2id"
)

func checkAppkey(bearerToken string) (models.User, models.AppkeyAuth, error) {
	var user models.User
	var keyAuth models.AppkeyAuth
	var appkey models.Appkey
	error := errors.New("Unauthorized")

	appkeyPrefix := strings.Split(bearerToken, ".")[0]
	if appkeyPrefix == "" {
		return user, keyAuth, error
	}

	appkey, err := cache.GetAppkey(appkeyPrefix)
	if err != nil {
		err = models.DB.Joins("User").Where("prefix = ?", appkeyPrefix).First(&appkey).Error
		if err != nil {
			return user, keyAuth, error
		}
	}
	go cacheAppkey(&appkey)

	match, err := argon2id.ComparePasswordAndHash(bearerToken, appkey.Hash)
	fmt.Println(match, err)
	if err != nil || !match {
		return user, keyAuth, error
	}

	user = appkey.User
	keyAuth.IsAppkey = true
	keyAuth.Permissions = appkey.Permissions

	return user, keyAuth, nil
}
