package session

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func setSessionCookies(c *gin.Context, sessionID uuid.UUID, sessionToken []byte) {
	createCookie(c, "session_id", sessionID.String())
	createCookie(c, "session_token", string(sessionToken))
}
