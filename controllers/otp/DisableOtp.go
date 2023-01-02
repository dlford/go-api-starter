package otp

import (
	"api/cache"
	"api/controllers/session"
	"api/models"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
)

// @Security Bearer
// @Summary Disable OTP
// @Description Allow sign in without OTP
// @Tags User Account
// @Accept json
// @Produce json
// @Param request body models.UserPasswordInput true "OTP code"
// @Success 200 {object} models.ResMsg
// @Failure 401 {object} models.ResErr
// @Failure 403 {object} models.ResErr
// @Failure 500 {object} models.ResErr
// @Router /user/otp [delete]
func DisableOtp(c *gin.Context) {
	var input models.UserPasswordInput
	err := models.ValidateInput(c, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResErr{Error: err.Error()})
		return
	}

	user, appkeyAuth, err := session.AuthorizeRequest(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResErr{Error: err.Error()})
		return
	}
	if appkeyAuth.IsAppkey {
		c.JSON(http.StatusForbidden, models.ResErr{Error: "Cannot disable OTP with appkey"})
		return
	}

	if !user.OtpEnabled {
		c.JSON(http.StatusBadRequest, models.ResErr{Error: "OTP is already disabled for this account"})
		return
	}

	var password models.Password
	if err := models.DB.Where("user_id = ?", user.ID).First(&password).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ResErr{Error: err.Error()})
		return
	}

	valid, err := argon2id.ComparePasswordAndHash(input.Password, password.Hash)
	if err != nil || !valid {
		c.JSON(http.StatusBadRequest, models.ResErr{Error: "Invalid password"})
		return
	}

	user.OtpEnabled = false
	err = models.DB.Save(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResErr{Error: err.Error()})
		return
	}
	cache.UpdateCachedUser(&user)

	models.DB.Where("user_id = ?", user.ID).Delete(&models.Otp{})
	models.DB.Where("user_id = ?", user.ID).Delete(&models.RecoveryCode{})

	c.JSON(http.StatusOK, models.ResMsg{Message: "OTP has been disabled"})
}
