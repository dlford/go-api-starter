package router

import (
	"api/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/swagger/index.html")
	})

	if constants.TRUSTED_PLATFORM != "" {
		r.TrustedPlatform = constants.TRUSTED_PLATFORM
	} else if len(constants.TRUSTED_PROXIES) > 0 {
		r.SetTrustedProxies(constants.TRUSTED_PROXIES)
	}

	swaggerRouter(r)
	userRouter(r)

	return r
}
