package user

import (
	"api/cache"
	"api/controllers/session"
	"api/models"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
)

// @Summary Reset password
// @Description Reset password using password reset link (omit `otp_code` if OTP is not enabled)
// @Tags Password
// @Accept json
// @Produce json
// @Param request body models.ResetUserPasswordInput true "Password Reset"
// @Success 200 {object} models.ResMsg
// @Failure 400 {object} models.ResErr
// @Failure 401 {object} models.ResErr
// @Failure 404 {object} models.ResErr
// @Failure 500 {object} models.ResErr
// @Router /user/password [post]
func ResetPassword(c *gin.Context) {
	var input models.ResetUserPasswordInput
	if err := models.ValidateInput(c, &input); err != nil {
		c.JSON(http.StatusBadRequest, models.ResErr{Error: err.Error()})
		return
	}

	var user models.User
	if err := models.DB.Where("id = ?", input.ID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ResErr{Error: "User not found"})
		return
	}

	token, err := cache.GetPasswordResetToken(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResErr{Error: "Password reset link expired"})
		return
	}

	if token != input.Token {
		c.JSON(http.StatusBadRequest, models.ResErr{Error: "Invalid password reset link"})
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

	if err := models.DB.Model(models.Password{}).Where("user_id = ?", user.ID).Update("hash", hash).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ResErr{Error: "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, models.ResMsg{Message: "Password reset"})
}
