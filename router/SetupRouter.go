package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	swaggerRouter(r)
	userRouter(r)

	return r
}
