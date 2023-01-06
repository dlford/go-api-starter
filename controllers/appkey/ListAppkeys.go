package appkey

import (
	"api/controllers/session"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security Bearer
// @Summary List application keys
// @Description List all application keys for the current user
// @Tags Application Keys
// @Accept json
// @Produce json
// @Success 200 {object} models.ResData[[]models.Appkey]
// @Failure 401 {object} models.ResErr
// @Failure 403 {object} models.ResErr
// @Failure 500 {object} models.ResErr
// @Router /user/appkeys [get]
func ListAppkeys(c *gin.Context) {
	user, appkeyAuth, err := session.AuthorizeRequest(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResErr{Error: err.Error()})
		return
	}
	if appkeyAuth.IsAppkey {
		c.JSON(http.StatusForbidden, models.ResErr{Error: "Cannot list appkeys with appkey"})
		return
	}

	var appkeys []models.Appkey
	if err := models.DB.Where("user_id = ?", user.ID).Find(&appkeys).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ResErr{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.ResData[[]models.Appkey]{Data: appkeys})
}
