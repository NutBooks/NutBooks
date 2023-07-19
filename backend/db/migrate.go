package db

import (
	"api/db/models"
	"log"
)

func MigrateMysql() {
	db, err := GetDB()

	err = db.AutoMigrate(
		&models.User{}, &models.Profile{}, &models.Authority{}, &models.Authentication{},
		&models.SocialLogin{}, &models.Password{}, &models.Cidi{},
		&models.Bookmark{}, &models.Keyword{},
	)
	if err != nil {
		log.Panic(err)
	}
}
