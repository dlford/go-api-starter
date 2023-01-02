package mail

import (
	"api/cache"
	"api/constants"
	"api/models"
	"net/url"

	"github.com/dchest/uniuri"
	"gopkg.in/gomail.v2"
)

func SendForgotPasswordEmail(user *models.User) {
	token := uniuri.NewLen(128)
	err := cache.AddPasswordResetToken(user, token)
	if err != nil {
		panic(err)
	}

	resetUrl, err := url.Parse(constants.RESET_PASSWORD_URL)
	if err != nil {
		panic(err)
	}
	params := resetUrl.Query()
	params.Add("id", url.QueryEscape(user.ID.String()))
	params.Add("token", url.QueryEscape(token))
	resetUrl.RawQuery = params.Encode()

	html, err := getHtmlFromTemplate("mail/templates/resetPassword.mjml", resetPasswordInput{
		FirstName: user.FirstName,
		ResetUrl:  resetUrl.String(),
	})
	if err != nil {
		panic(err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(constants.MAIL_FROM, constants.MAIL_FROM_NAME))
	m.SetHeader("To", m.FormatAddress(user.Email, user.FirstName+" "+user.LastName))
	m.SetHeader("Subject", "Reset your password")
	m.SetBody("text/plain", "Reset your password by clicking the link below (this link will expire in 15 minutes):\n\n"+resetUrl.String())
	m.AddAlternative("text/html", html)

	d := gomail.NewDialer(constants.MAIL_HOST, constants.MAIL_PORT, constants.MAIL_USER, constants.MAIL_PASSWORD)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
