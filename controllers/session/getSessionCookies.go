package session

import (
	"api/models"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type SessionCookies struct {
	sessionId    string
	sessionToken string
}

func getSessionCookies(c *gin.Context) (SessionCookies, error) {
	sessionIdCookie, err := c.Request.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResErr{Error: "Unauthorized"})
		return SessionCookies{}, err
	}

	sessionTokenCookie, err := c.Request.Cookie("session_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResErr{Error: "Unauthorized"})
		return SessionCookies{}, err
	}

	// NOTE: issue for cookie values being URL encoded - https://github.com/gin-gonic/gin/issues/1717
	sessionId, _ := url.QueryUnescape(sessionIdCookie.Value)
	sessionToken, _ := url.QueryUnescape(sessionTokenCookie.Value)

	return SessionCookies{
		sessionId,
		sessionToken,
	}, nil
}
