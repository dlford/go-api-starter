package user

import (
	"api/cache"
	"api/controllers/session"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security Bearer
// @Summary Delete account
// @Description Delete user account and all data (omit `otp_code` if OTP is not enabled, omit `password` if OTP is enabled)
// @Tags User Account
// @Accept json
// @Produce json
// @Param request body models.DeleteUserInput true "Credentials"
// @Success 200 {object} models.ResMsg
// @Failure 400 {object} models.ResErr
// @Failure 401 {object} models.ResErr
// @Failure 403 {object} models.ResErr
// @Failure 500 {object} models.ResErr
// @Router /user [delete]
func DeleteUser(c *gin.Context) {
	var input models.DeleteUserInput
	if err := models.ValidateInput(c, &input); err != nil {
		c.JSON(http.StatusBadRequest, models.ResErr{Error: err.Error()})
		return
	}

	user, appkeyAuth, err := session.AuthorizeRequest(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResErr{Error: err.Error()})
		return
	}
	if appkeyAuth.IsAppkey {
		c.JSON(http.StatusForbidden, models.ResErr{Error: "Cannot delete account with appkey"})
		return
	}

	if err := session.ValidatePasswordOrOtpCode(user, input.Password, input.OtpCode); err != nil {
		c.JSON(http.StatusUnauthorized, models.ResErr{Error: err.Error()})
		return
	}

	models.DB.Model(&models.User{}).Delete(&user)

	session.ClearSessionCookies(c)
	cache.RemoveUserFromCache(&user)

	c.JSON(http.StatusNotImplemented, models.ResMsg{Message: "User and data have been removed"})
}
