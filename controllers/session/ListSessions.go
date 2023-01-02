package session

import (
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security Bearer
// @Summary List sessions
// @Description List all sessions for the current logged in user
// @Tags User Account
// @Accept json
// @Produce json
// @Success 200 {object} models.ResData[[]models.Session]
// @Failure 401 {object} models.ResErr
// @Failure 403 {object} models.ResErr
// @Failure 500 {object} models.ResErr
// @Router /user/sessions [get]
func ListSessions(c *gin.Context) {
	user, appkeyAuth, err := AuthorizeRequest(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResErr{Error: err.Error()})
		return
	}
	if appkeyAuth.IsAppkey {
		c.JSON(http.StatusForbidden, models.ResErr{Error: "Cannot list sessions with appkey"})
		return
	}

	var sessions []models.Session
	err = models.DB.Where("user_id = ?", user.ID).Find(&sessions).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResErr{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.ResData[[]models.Session]{Data: sessions})
}
