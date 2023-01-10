package appkey

import (
	"api/controllers/session"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// @Security Bearer
// @Summary Delete application key
// @Description Delete an application key for the current user
// @Tags Application Keys
// @Accept json
// @Produce json
// @Param appkey_id path string true "Appkey ID" Format(uuid)
// @Success 200 {object} models.ResMsg
// @Failure 401 {object} models.ResErr
// @Failure 403 {object} models.ResErr
// @Failure 404 {object} models.ResErr
// @Failure 500 {object} models.ResErr
// @Router /user/appkeys/{appkey_id} [delete]
func DeleteAppkey(c *gin.Context) {
	appkeyId, err := uuid.FromString(c.Param("appkey_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResErr{Error: err.Error()})
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
