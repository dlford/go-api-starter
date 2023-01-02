package session

import (
	"api/constants"

	"github.com/gin-gonic/gin"
)

func createCookie(c *gin.Context, name string, value string) {
	maxAge := constants.SESSION_LIFETIME_SECONDS
	if value == "" {
		maxAge = 0
	}

	c.SetSameSite(constants.COOKIE_SAME_SITE)
	c.SetCookie(
		name,
		value,
		maxAge,
		constants.COOKIE_PATH,
		constants.COOKIE_DOMAIN,
		constants.COOKIE_SECURE,
		constants.COOKIE_HTTP_ONLY,
	)
}
