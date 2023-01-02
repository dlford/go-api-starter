package main

import (
	"api/cache"
	"api/constants"
	"api/jobs"
	"api/migrations"
	"api/models"
	"api/router"
)

func main() {
	constants.SetupEnv()

	cache.ConnectCache()

	models.ConnectDatabase()
	migrations.MigrateDatabase()

	jobs.StartJobs()

	r := router.SetupRouter()
	r.Run()
}
