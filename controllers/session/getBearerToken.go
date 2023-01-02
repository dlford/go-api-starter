package session

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func getBearerToken(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")

	authFields := strings.Fields(authHeader)
	if len(authFields) != 2 || strings.ToLower(authFields[0]) != "bearer" {
		return "", errors.New("Unauthorized")
	}

	bearerToken := authFields[1]

	if bearerToken == "" {
		return "", errors.New("Unauthorized")
	}

	return bearerToken, nil
}
