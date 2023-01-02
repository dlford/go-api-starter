package appkey

import (
	"api/controllers/session"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security Bearer
// @Summary Delete application key
// @Description Delete an application key for the current user
// @Tags User Account
// @Accept json
// @Produce json
// @Param id path string true "Appkey ID" Format(uuid)
// @Success 200 {object} models.ResMsg
// @Failure 401 {object} models.ResErr
// @Failure 403 {object} models.ResErr
// @Failure 404 {object} models.ResErr
// @Failure 500 {object} models.ResErr
// @Router /user/appkeys/{id} [delete]
func DeleteAppkey(c *gin.Context) {
	appkeyId := c.Param("id")
	if appkeyId == "" {
		c.JSON(http.StatusBadRequest, models.ResErr{Error: "Appkey ID required in path"})
		return
	}

	user, appkeyAuth, err := session.AuthorizeRequest(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResErr{Error: err.Error()})
		return
	}
	if appkeyAuth.IsAppkey {
		c.JSON(http.StatusForbidden, models.ResErr{Error: "Cannot delete appkey with appkey"})
		return
	}

	tx := models.DB.Where("id = ? AND user_id = ?", appkeyId, user.ID).Delete(&models.Appkey{})
	if tx.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, models.ResErr{Error: "Appkey not found"})
		return
	}

	c.JSON(http.StatusOK, models.ResMsg{Message: "Appkey deleted"})
}
