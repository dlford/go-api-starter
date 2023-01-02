package mail

import (
	"net/url"

	"github.com/dchest/uniuri"
	"gopkg.in/gomail.v2"

	"api/cache"
	"api/constants"
	"api/models"
)

func SendEmailVerification(user *models.User) {
	token := uniuri.NewLen(128)
	err := cache.AddEmailVerificationToken(user, token)
	if err != nil {
		panic(err)
	}

	verifyUrl, err := url.Parse(constants.FQDN)
	if err != nil {
		panic(err)
	}
	verifyUrl.Path += "/user/email"
	params := verifyUrl.Query()
	params.Add("id", url.QueryEscape(user.ID.String()))
	params.Add("token", url.QueryEscape(token))
	if constants.EMAIL_VERIFIED_REDIRECT_URL != "" {
		params.Add("redirect_url", url.QueryEscape(constants.EMAIL_VERIFIED_REDIRECT_URL))
	}
	verifyUrl.RawQuery = params.Encode()

	html, err := getHtmlFromTemplate("mail/templates/emailVerification.mjml", emailVerificationInput{
		FirstName: user.FirstName,
		VerifyUrl: verifyUrl.String(),
	})
	if err != nil {
		panic(err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(constants.MAIL_FROM, constants.MAIL_FROM_NAME))
	m.SetHeader("To", m.FormatAddress(user.Email, user.FirstName+" "+user.LastName))
	m.SetHeader("Subject", "Verify your email address")
	m.SetBody("text/plain", "Please verify your email address by clicking the link below (this link will expire in 7 days):\n\n"+verifyUrl.String())
	m.AddAlternative("text/html", html)

	d := gomail.NewDialer(constants.MAIL_HOST, constants.MAIL_PORT, constants.MAIL_USER, constants.MAIL_PASSWORD)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
