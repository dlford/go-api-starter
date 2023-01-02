package session

import (
	"api/cache"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security Bearer
// @Summary Sign out
// @Description Delete current session and invalidate its bearer tokens
// @Tags User Account
// @Accept json
// @Produce json
// @Success 200 {object} models.ResMsg
// @Failure 401 {object} models.ResErr
// @Failure 403 {object} models.ResErr
// @Failure 500 {object} models.ResErr
// @Router /user/sessions [delete]
func SignOut(c *gin.Context) {
	user, appkeyAuth, err := AuthorizeRequest(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResErr{Error: err.Error()})
		return
	}
	if appkeyAuth.IsAppkey {
		c.JSON(http.StatusForbidden, models.ResErr{Error: "Cannot sign out with appkey"})
		return
	}

	cookies, err := getSessionCookies(c)
	if err == nil {
		models.DB.Where("id = ?", cookies.sessionId).Delete(&models.Session{})
		cache.RemoveSessionFromCache(user.ID.String(), cookies.sessionId)
	}

	ClearSessionCookies(c)

	c.JSON(http.StatusOK, models.ResMsg{Message: "Signed out"})
}
