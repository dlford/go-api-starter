package migrations

import "api/models"

func MigrateDatabase() {
	err := models.DB.AutoMigrate(
		&models.User{},
		&models.Session{},
		&models.Password{},
		&models.Otp{},
		&models.RecoveryCode{},
		&models.Appkey{},
	)

	if err != nil {
		panic("Failed to migrate database!")
	}
}
