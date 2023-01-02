package user

import (
	"net/http"
	"regexp"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"api/controllers/session"
	"api/mail"
	"api/models"
)

// @Summary Sign up
// @Description Create an account
// @Tags User Account
// @Accept json
// @Produce json
// @Param request body models.CreateUserInput true "User credentials"
// @Success 200 {object} models.ResData[models.BearerTokenResponse]
// @Failure 400 {object} models.ResErr
// @Failure 401 {object} models.ResErr
// @Failure 404 {object} models.ResErr
// @Failure 500 {object} models.ResErr
// @Router /user [post]
func CreateUser(c *gin.Context) {
	var input models.CreateUserInput
	if err := models.ValidateInput(c, &input); err != nil {
		c.JSON(http.StatusBadRequest, models.ResErr{Error: err.Error()})
		return
	}

	hash, err := argon2id.CreateHash(input.Password, argon2id.DefaultParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResErr{Error: err.Error()})
		return
	}

	user := models.User{
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Password: models.Password{
			Hash: hash,
		},
	}

	err = models.DB.Create(&user).Error
	if err != nil {
		match, _ := regexp.MatchString("duplicate key value violates unique constraint.*users_email_key", err.Error())
		if match {
			c.JSON(http.StatusUnprocessableEntity, models.ResErr{Error: "An account already exists for " + input.Email})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ResErr{Error: err.Error()})
		return
	}

	go mail.SendEmailVerification(&user)

	session.CreateOrRenewSession(c, &user, uuid.Nil)
}
