package session

import (
	"github.com/gin-gonic/gin"
)

func ClearSessionCookies(c *gin.Context) {
	createCookie(c, "session_id", "")
	createCookie(c, "session_token", "")
}
