package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"api/controllers/session"
	"api/models"
)

// @Security Bearer
// @Summary Get current user
// @Description Return the current logged in user details
// @Tags User Account
// @Accept json
// @Produce json
// @Success 200 {object} models.ResData[models.User]
// @Failure 401 {object} models.ResErr
// @Failure 403 {object} models.ResErr
// @Router /user [get]
func GetCurrentUser(c *gin.Context) {
	user, appkeyAuth, err := session.AuthorizeRequest(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResErr{Error: err.Error()})
		return
	}
	if appkeyAuth.IsAppkey && !appkeyAuth.Permissions.CanViewUserInfo {
		c.JSON(http.StatusForbidden, models.ResErr{Error: "Permission denied"})
		return
	}

	c.JSON(http.StatusOK, models.ResData[models.User]{Data: user})
}
