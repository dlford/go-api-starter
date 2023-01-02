package jobs

import (
	"api/constants"
	"api/models"
	"time"

	"github.com/go-co-op/gocron"
)

func sessionReaper() {
	s := gocron.NewScheduler(time.Local)
	s.Cron(constants.SESSION_REAPER_SCHEDULE).Do(reapSessions)

	s.StartAsync()
}

func reapSessions() {
	models.DB.Exec("DELETE FROM sessions WHERE expires_at < NOW()")
}
