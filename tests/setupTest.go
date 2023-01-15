package tests

import (
	"api/cache"
	"api/constants"
	"api/migrations"
	"api/models"
	"api/router"

	"github.com/gin-gonic/gin"
)

func setupTest() *gin.Engine {
	gin.SetMode(gin.TestMode)
	constants.SetupEnv()
	cache.ConnectCache()
	models.ConnectDatabase()
	migrations.MigrateDatabase()
	return router.SetupRouter()
}
