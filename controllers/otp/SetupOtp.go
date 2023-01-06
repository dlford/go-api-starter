package otp

import (
	"api/constants"
	"api/controllers/session"
	"api/models"
	"bytes"
	"encoding/base64"
	"image/png"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp/totp"
)

// @Security Bearer
// @Summary Set up OTP
// @Description Generate a QR code for setting up OTP
// @Tags Two Factor Authentication
// @Accept json
// @Produce json
// @Success 200 {object} models.ResData[models.SetupOtpResponse]
// @Failure 401 {object} models.ResErr
// @Failure 403 {object} models.ResErr
// @Failure 500 {object} models.ResErr
// @Router /user/otp [get]
func SetupOtp(c *gin.Context) {
	user, appkeyAuth, err := session.AuthorizeRequest(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResErr{Error: err.Error()})
		return
	}
	if appkeyAuth.IsAppkey {
		c.JSON(http.StatusForbidden, models.ResErr{Error: "Cannot set up OTP for appkey"})
		return
	}

	if user.OtpEnabled {
		c.JSON(http.StatusBadRequest, models.ResErr{Error: "OTP is already enabled for this account"})
		return
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      constants.SERVICE_NAME,
		AccountName: user.Email,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResErr{Error: err.Error()})
		return
	}

	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResErr{Error: err.Error()})
		return
	}
	png.Encode(&buf, img)
	imgStr := "data:image/png;base64,"
	imgStr += base64.StdEncoding.EncodeToString(buf.Bytes())

	var otp models.Otp
	models.DB.Model(&user).Association("Otp").Find(&otp)
	otp.Secret = key.Secret()

	err = models.DB.Save(&otp).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResErr{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.SetupOtpResponse{
		Secret: key.Secret(),
		QrCode: imgStr,
	})
}
