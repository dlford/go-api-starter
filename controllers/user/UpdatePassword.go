package user

import (
	"api/controllers/session"
	"api/models"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
)

// @Security Bearer
// @Summary Change password
// @Description Change password using existing password (omit `otp_code` if OTP is not enabled)
// @Tags User Account
// @Accept json
// @Produce json
// @Param request body models.UpdateUserPasswordInput true "Password Reset"
// @Success 200 {object} models.ResMsg
// @Failure 400 {object} models.ResErr
// @Failure 401 {object} models.ResErr
// @Failure 403 {object} models.ResErr
// @Failure 500 {object} models.ResErr
// @Router /user/password [put]
func UpdatePassword(c *gin.Context) {
	var input models.UpdateUserPasswordInput
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
		c.JSON(http.StatusForbidden, models.ResErr{Error: "Cannot change password with appkey"})
		return
	}

	var password models.Password
	if err := models.DB.Where("user_id = ?", user.ID).First(&password).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ResErr{Error: "Internal server error"})
		return
	}

	match, err := argon2id.ComparePasswordAndHash(input.Password, password.Hash)
	if err != nil || !match {
		c.JSON(http.StatusUnauthorized, models.ResErr{Error: "Invalid password"})
		return
	}

	if err := session.ValidateOtpCode(user, input.OtpCode); err != nil {
		c.JSON(http.StatusUnauthorized, models.ResErr{Error: err.Error()})
		return
	}

	hash, err := argon2id.CreateHash(input.NewPassword, argon2id.DefaultParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResErr{Error: "Internal server error"})
		return
	}

	password.Hash = hash
	if err := models.DB.Save(&password).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ResErr{Error: "Internal server error"})
		return
	}

	c.JSON(http.StatusNotImplemented, models.ResMsg{Message: "Password updated"})
}
