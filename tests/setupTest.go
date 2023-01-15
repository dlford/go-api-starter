package tests

import (
	"api/cache"
	"api/constants"
	"api/migrations"
	"api/models"
	"api/router"
	"os"
	"path"
	"runtime"

	"github.com/gin-gonic/gin"
)

func setupTest() *gin.Engine {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}

	gin.SetMode(gin.TestMode)
	constants.SetupEnv()
	cache.ConnectCache()
	models.ConnectDatabase()
	migrations.MigrateDatabase()
	return router.SetupRouter()
}
