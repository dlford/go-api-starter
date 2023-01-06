package user

import (
	"api/cache"
	"api/models"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// @Summary Verify email address
// @Description Verify the user's email address, redirects to `redirect_url` if provided and appends the query parameters `email_verified`[boolean] and `error`[string]
// @Tags Email Address
// @Accept json
// @Produce json
// @Param id query string true "User ID" Format(uuid)
// @param token query string true "Verification Token"
// @param redirect_url query string false "Redirect URL"
// @Success 200 {object} models.ResMsg
// @Failure 400 {object} models.ResErr
// @Failure 500 {object} models.ResErr
// @Router /user/email [get]
func VerifyEmailAddress(c *gin.Context) {
	var input models.VerifyEmailAddressInput
	err := c.ShouldBindQuery(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResErr{Error: err.Error()})
		return
	}

	hasRedirectUrl := input.RedirectUrl != ""
	redirectUrlStr, err := url.QueryUnescape(input.RedirectUrl)
	if err != nil {
		hasRedirectUrl = false
		fmt.Println(err)
	}
	redirectUrl, err := url.Parse(redirectUrlStr)
	if err != nil {
		hasRedirectUrl = false
		fmt.Println(err)
	}
	params := redirectUrl.Query()

	var user models.User
	if err := models.DB.Where("id = ?", input.ID).First(&user).Error; err != nil {
		if !hasRedirectUrl {
			c.JSON(http.StatusBadRequest, models.ResErr{Error: "Record not found"})
			return
		}
		params.Add("email_verified", "false")
		params.Add("error", "record_not_found")
		redirectUrl.RawQuery = params.Encode()
		c.Redirect(http.StatusTemporaryRedirect, redirectUrl.String())
		return
	}

	if user.EmailVerified {
		if !hasRedirectUrl {
			c.JSON(http.StatusBadRequest, models.ResErr{Error: "Email already verified"})
			return
		}
		params.Add("email_verified", "true")
		params.Add("error", "email_already_verified")
		redirectUrl.RawQuery = params.Encode()
		c.Redirect(http.StatusTemporaryRedirect, redirectUrl.String())
		return
	}

	code, err := cache.GetEmailVerificationToken(&user)
	if err != nil {
		if !hasRedirectUrl {
			c.JSON(http.StatusBadRequest, models.ResErr{Error: "Invalid token"})
			return
		}
		params.Add("email_verified", "false")
		params.Add("error", "invalid_token")
		redirectUrl.RawQuery = params.Encode()
		c.Redirect(http.StatusTemporaryRedirect, redirectUrl.String())
		return
	}

	if code != input.Token {
		fmt.Println(err)
		if !hasRedirectUrl {
			c.JSON(http.StatusBadRequest, models.ResErr{Error: "Invalid token"})
			return
		}
		params.Add("email_verified", "false")
		params.Add("error", "invalid_token")
		redirectUrl.RawQuery = params.Encode()
		c.Redirect(http.StatusTemporaryRedirect, redirectUrl.String())
		return
	}

	user.EmailVerified = true
	models.DB.Save(&user)
	cache.UpdateCachedUser(&user)

	go cache.RemoveEmailVerificationToken(&user)

	if !hasRedirectUrl {
		c.JSON(http.StatusOK, models.ResMsg{Message: "Email verified"})
		return
	}
	params.Add("email_verified", "true")
	redirectUrl.RawQuery = params.Encode()
	c.Redirect(http.StatusTemporaryRedirect, redirectUrl.String())
}
