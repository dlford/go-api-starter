package router

import (
	"api/constants"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	if constants.TRUSTED_PLATFORM != "" {
		r.TrustedPlatform = constants.TRUSTED_PLATFORM
	} else if len(constants.TRUSTED_PROXIES) > 0 {
		r.SetTrustedProxies(constants.TRUSTED_PROXIES)
	}

	swaggerRouter(r)
	userRouter(r)

	return r
}
