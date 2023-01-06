package user

import (
	"api/controllers/session"
	"api/mail"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security Bearer
// @Summary Retry email verification
// @Description Re-send the email verification email (link is valid for 7 days)
// @Tags Email Address
// @Accept json
// @Produce json
// @Success 200 {object} models.ResMsg
// @Failure 400 {object} models.ResErr
// @Failure 401 {object} models.ResErr
// @Failure 403 {object} models.ResErr
// @Router /user/email [patch]
func ResendEmailVerification(c *gin.Context) {
	user, appkeyAuth, err := session.AuthorizeRequest(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResErr{Error: err.Error()})
		return
	}
	if appkeyAuth.IsAppkey {
		c.JSON(http.StatusForbidden, models.ResErr{Error: "Cannot resend email verification with appkey"})
		return
	}

	if user.EmailVerified {
		c.JSON(http.StatusBadRequest, models.ResErr{
			Error: "Email already verified",
		})
		return
	}

	go mail.SendEmailVerification(&user)

	c.JSON(http.StatusOK, models.ResMsg{
		Message: "Verification email sent",
	})
}
