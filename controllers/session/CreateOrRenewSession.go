package session

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/dchest/uniuri"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"api/cache"
	"api/constants"
	"api/models"
)

func CreateOrRenewSession(c *gin.Context, user *models.User, sessionId uuid.UUID) {
	expiresAt := time.Now().AddDate(0, 0, 30)
	salt := uniuri.NewLen(16)

	var session models.Session

	if sessionId == uuid.Nil {
		session = models.Session{
			UserId:    user.ID,
			ExpiresAt: expiresAt,
			Salt:      salt,
			UserAgent: c.Request.UserAgent(),
			IpAddress: c.ClientIP(),
		}
		models.DB.Create(&session)
	} else {
		models.DB.Where("id = ?", sessionId).First(&session)
		session.Salt = salt
		session.ExpiresAt = expiresAt
		session.UserAgent = c.Request.UserAgent()
		session.IpAddress = c.ClientIP()
		models.DB.Save(&session)
	}

	sessionToken, err := argon2id.CreateHash(salt+"."+session.ID.String(), argon2id.DefaultParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResErr{Error: err.Error()})
		return
	}

	setSessionCookies(c, session.ID, []byte(sessionToken))

	accessToken := uniuri.NewLen(128)

	err = cache.AddAccessToken(user, session.ID, accessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResErr{Error: err.Error()})
		return
	}

	c.Header("X-Bearer-Token-TTL", fmt.Sprint(constants.ACCESS_TOKEN_LIFETIME.Seconds()))
	c.JSON(http.StatusOK, models.ResData[models.BearerTokenResponse]{
		Data: models.BearerTokenResponse{
			BearerToken: "Bearer " + accessToken,
		},
	})
}
