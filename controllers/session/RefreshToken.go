package session

import (
	"api/models"
	"net/http"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Refresh token
// @Description Get a new bearer token using session cookies
// @Tags User Account
// @Accept json
// @Produce json
// @Success 200 {object} models.ResData[models.BearerTokenResponse]
// @Failure 401 {object} models.ResErr
// @Failure 500 {object} models.ResErr
// @Router /user/sessions [patch]
func RefreshToken(c *gin.Context) {
	cookies, err := getSessionCookies(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResErr{Error: "Unauthorized"})
		return
	}

	var session models.Session
	err = models.DB.Model(&models.Session{}).Find(&session, "id = ?", cookies.sessionId).Error
	if err != nil || session.ID == uuid.Nil {
		c.JSON(http.StatusUnauthorized, models.ResErr{Error: "Unauthorized"})
		return
	}
	if session.ExpiresAt.Before(time.Now()) {
		models.DB.Delete(&session)
		c.JSON(http.StatusUnauthorized, models.ResErr{Error: "Session expired, please log in again"})
		return
	}

	match, err := argon2id.ComparePasswordAndHash(session.Salt+"."+session.ID.String(), cookies.sessionToken)
	if err != nil || !match {
		c.JSON(http.StatusUnauthorized, models.ResErr{Error: "Unauthorized"})
		return
	}

	var user models.User
	err = models.DB.Model(&models.User{}).Find(&user, "id = ?", session.UserId).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResErr{Error: "Unauthorized"})
		return
	}

	CreateOrRenewSession(c, &user, session.ID)
}
