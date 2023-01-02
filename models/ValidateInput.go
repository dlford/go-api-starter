package models

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func ValidateInput(c *gin.Context, input interface{}) error {
	if err := c.ShouldBindJSON(&input); err != nil {
		if err.Error() == "EOF" {
			return errors.New("No JSON input provided")
		}
		return err
	}
	return nil
}
