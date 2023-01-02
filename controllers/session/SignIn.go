package session

import (
	"api/models"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Sign In
// @Description Create a new session (omit `otp_code` if OTP is not enabled)
// @Tags User Account
// @Accept json
// @Produce json
// @Param request body models.SignInInput true "User credentials"
// @Success 200 {object} models.ResData[models.BearerTokenResponse]
// @Failure 400 {object} models.ResErr
// @Failure 401 {object} models.ResErr
// @Failure 500 {object} models.ResErr
// @Router /user/sessions [post]
func SignIn(c *gin.Context) {
	var input models.SignInInput
	if err := models.ValidateInput(c, &input); err != nil {
		c.JSON(http.StatusBadRequest, models.ResErr{Error: err.Error()})
		return
	}

	var user models.User
	err := models.DB.Joins("Password").Find(&user, "email = ?", input.Email).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResErr{Error: "Invalid email or password"})
		return
	}

	match, err := argon2id.ComparePasswordAndHash(input.Password, user.Password.Hash)
	if err != nil || !match {
		c.JSON(http.StatusBadRequest, models.ResErr{Error: "Invalid email or password"})
		return
	}

	if err := ValidateOtpCode(user, input.OtpCode); err != nil {
		c.JSON(http.StatusUnauthorized, models.ResErr{Error: err.Error()})
		return
	}

	CreateOrRenewSession(c, &user, uuid.Nil)
}
