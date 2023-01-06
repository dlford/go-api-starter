package otp

import (
	"api/cache"
	"api/controllers/session"
	"api/models"
	"net/http"

	"github.com/dchest/uniuri"
	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp/totp"
)

// @Security Bearer
// @Summary Enable OTP
// @Description Prevent logging in without OTP
// @Tags Two Factor Authentication
// @Accept json
// @Produce json
// @Param request body models.EnableOtpInput true "OTP code"
// @Success 200 {object} models.ResData[models.EnableOtpResponse]
// @Failure 401 {object} models.ResErr
// @Failure 403 {object} models.ResErr
// @Failure 500 {object} models.ResErr
// @Router /user/otp [patch]
func EnableOtp(c *gin.Context) {
	var input models.EnableOtpInput
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
		c.JSON(http.StatusForbidden, models.ResErr{Error: "Cannot enable OTP with appkey"})
		return
	}

	if user.OtpEnabled {
		c.JSON(http.StatusBadRequest, models.ResErr{Error: "OTP is already enabled for this account"})
		return
	}

	var otp models.Otp
	models.DB.Model(&user).Association("Otp").Find(&otp)

	if otp.Secret == "" {
		c.JSON(http.StatusBadRequest, models.ResErr{Error: "OTP has not been set up for this account"})
		return
	}

	valid := totp.Validate(input.OtpCode, otp.Secret)
	if !valid {
		c.JSON(http.StatusBadRequest, models.ResErr{Error: "Invalid OTP code"})
		return
	}

	user.OtpEnabled = true
	err = models.DB.Save(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResErr{Error: err.Error()})
		return
	}
	cache.UpdateCachedUser(&user)

	var backupCodes [10]string
	for i := range backupCodes {
		code := uniuri.NewLen(12)
		backupCodes[i] = code
		recoveryCode := models.RecoveryCode{
			UserId: user.ID,
			Code:   code,
		}
		if err := models.DB.Create(&recoveryCode).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.ResErr{Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, models.ResData[models.EnableOtpResponse]{
		Data: models.EnableOtpResponse{
			RecoveryCodes: backupCodes,
		},
	})
}
