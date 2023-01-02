package user

import (
	"api/mail"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Forgot password
// @Description Request a password reset email (link is valid for 15 minutes)
// @Tags User Account
// @Accept json
// @Produce json
// @Param request body models.UserEmailInput true "Email"
// @Success 200 {object} models.ResMsg
// @Failure 400 {object} models.ResErr
// @Router /user/password [patch]
func ForgotPassword(c *gin.Context) {
	var input models.UserEmailInput
	if err := models.ValidateInput(c, &input); err != nil {
		c.JSON(http.StatusBadRequest, models.ResErr{Error: err.Error()})
		return
	}

	var user models.User
	if err := models.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		// Don't tell the user that the email doesn't exist
		c.JSON(http.StatusOK, models.ResMsg{Message: "Email sent"})
		return
	}

	go mail.SendForgotPasswordEmail(&user)

	c.JSON(http.StatusOK, models.ResMsg{Message: "Email sent"})
}
