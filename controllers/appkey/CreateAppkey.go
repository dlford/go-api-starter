package appkey

import (
	"api/controllers/session"
	"api/models"
	"net/http"
	"reflect"

	"github.com/alexedwards/argon2id"
	"github.com/dchest/uniuri"
	"github.com/gin-gonic/gin"
)

// @Security Bearer
// @Summary Create application key
// @Description Create an application key for the current user (omit `otp_code` if OTP is not enabled, omit `password` if OTP is enabled)
// @Tags User Account
// @Accept json
// @Produce json
// @Param input body models.CreateAppkeyInput true "Name and permissions"
// @Success 200 {object} models.ResData[models.CreateAppkeyResponse]
// @Failure 401 {object} models.ResErr
// @Failure 403 {object} models.ResErr
// @Failure 500 {object} models.ResErr
// @Router /user/appkeys [post]
func CreateAppkey(c *gin.Context) {
	var input models.CreateAppkeyInput
	err := models.ValidateInput(c, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResErr{Error: err.Error()})
	}

	values := reflect.ValueOf(input.Permissions)
	noPermissions := true
	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).Bool() {
			noPermissions = false
			break
		}
	}
	if noPermissions {
		c.JSON(http.StatusBadRequest, models.ResErr{Error: "Cannot create appkey with no permissions"})
		return
	}

	user, appkeyAuth, err := session.AuthorizeRequest(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResErr{Error: err.Error()})
		return
	}
	if appkeyAuth.IsAppkey {
		c.JSON(http.StatusForbidden, models.ResErr{Error: "Cannot create appkey with appkey"})
		return
	}
	if !user.EmailVerified {
		c.JSON(http.StatusForbidden, models.ResErr{Error: "Please verify your email address"})
	}

	if err := session.ValidatePasswordOrOtpCode(user, input.Password, input.OtpCode); err != nil {
		c.JSON(http.StatusUnauthorized, models.ResErr{Error: err.Error()})
		return
	}

	prefix := uniuri.NewLen(16)
	secret := uniuri.NewLen(128)
	key := prefix + "." + secret
	hash, err := argon2id.CreateHash(key, argon2id.DefaultParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResErr{Error: err.Error()})
		return
	}

	appkey := models.Appkey{
		UserId:      user.ID,
		Name:        input.Name,
		Prefix:      prefix,
		Hash:        hash,
		Permissions: input.Permissions,
	}

	if err := models.DB.Create(&appkey).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ResErr{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.ResData[models.CreateAppkeyResponse]{
		Data: models.CreateAppkeyResponse{
			Key: "Bearer " + key,
		},
	})
}
