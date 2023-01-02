package user

import (
	"api/cache"
	"api/controllers/session"
	"api/mail"
	"api/models"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

// @Security Bearer
// @Summary Update user
// @Description Change user account information
// @Tags User Account
// @Accept json
// @Produce json
// @Param request body models.UpdateUserInput true "User details"
// @Success 200 {object} models.ResData[models.User]
// @Failure 400 {object} models.ResErr
// @Failure 401 {object} models.ResErr
// @Failure 403 {object} models.ResErr
// @Failure 500 {object} models.ResErr
// @Router /user [put]
func UpdateUser(c *gin.Context) {
	var input models.UpdateUserInput
	if err := models.ValidateInput(c, &input); err != nil {
		c.JSON(http.StatusBadRequest, models.ResErr{Error: err.Error()})
		return
	}

	user, appkeyAuth, err := session.AuthorizeRequest(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResErr{Error: err.Error()})
		return
	}
	if appkeyAuth.IsAppkey && !appkeyAuth.Permissions.CanEditUserInfo {
		c.JSON(http.StatusForbidden, models.ResErr{Error: "Permission denied"})
		return
	}

	err = models.DB.Model(models.User{}).Where("id = ?", user.ID).First(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResErr{
			Error: "Internal server error",
		})
		return
	}

	emailChanged := false
	if input.Email != user.Email {
		user.EmailVerified = false
		emailChanged = true
	}

	user.Email = input.Email
	user.FirstName = input.FirstName
	user.LastName = input.LastName

	if err := models.DB.Save(&user).Error; err != nil {
		match, _ := regexp.MatchString("duplicate key value violates unique constraint.*users_email_key", err.Error())
		if match {
			c.JSON(http.StatusUnprocessableEntity, models.ResErr{Error: "An account already exists for " + input.Email})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ResErr{
			Error: "Internal server error",
		})
		return
	}

	if emailChanged {
		go mail.SendEmailVerification(&user)
	}

	go cache.UpdateCachedUser(&user)

	c.JSON(http.StatusOK, models.ResData[models.User]{
		Data: user,
	})
}
