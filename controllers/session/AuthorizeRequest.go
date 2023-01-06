package session

import (
	"api/cache"
	"api/models"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthorizeRequest(c *gin.Context) (models.User, models.AppkeyAuth, error) {
	var user models.User
	var keyAuth models.AppkeyAuth
	error := errors.New("Unauthorized")

	bearerToken, err := getBearerToken(c)
	if err != nil {
		return user, keyAuth, error
	}

	user, ttl, err := cache.GetUserFromAccessToken(bearerToken)
	if err != nil {
		user, keyAuth, err = checkAppkey(bearerToken)
		if err != nil {
			return user, keyAuth, error
		}
	}

	c.Header("X-Bearer-Token-TTL", fmt.Sprint(ttl))
	return user, keyAuth, nil
}
