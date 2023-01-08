package session

import (
	"api/cache"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Security Bearer
// @Summary Revoke session
// @Description Revoke access to a session and invalidate its bearer tokens
// @Tags Sessions
// @Accept json
// @Produce json
// @Param sessionId path string true "Session ID" Format(uuid)
// @Success 200 {object} models.ResMsg
// @Failure 401 {object} models.ResErr
// @Failure 403 {object} models.ResErr
// @Failure 404 {object} models.ResErr
// @Failure 500 {object} models.ResErr
// @Router /user/sessions/{sessionId} [delete]
func DeleteSession(c *gin.Context) {
	sessionId, err := uuid.Parse(c.Param("sessionId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResErr{Error: err.Error()})
		return
	}

	user, appkeyAuth, err := AuthorizeRequest(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResErr{Error: err.Error()})
		return
	}
	if appkeyAuth.IsAppkey {
		c.JSON(http.StatusForbidden, models.ResErr{Error: "Cannot revoke session with appkey"})
		return
	}

	cache.RemoveSessionFromCache(user.ID.String(), sessionId.String())
	tx := models.DB.Where("id = ? AND user_id = ?", sessionId, user.ID).Delete(&models.Session{})
	if tx.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, models.ResErr{Error: "Session not found"})
		return
	}

	c.JSON(http.StatusOK, models.ResMsg{Message: "Session signed out"})
}
