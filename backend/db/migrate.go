package db

import (
	"api/db/models"
	"log"
)

func MigrateMysql() {
	err := DB.AutoMigrate(
		&models.User{}, &models.Profile{}, &models.Authority{}, &models.Authentication{},
		&models.SocialLogin{}, &models.Password{}, &models.Cidi{},
		&models.Bookmark{}, &models.Keyword{},
	)
	if err != nil {
		log.Panicf("Cannot migrate DB: %v", err)
	}

	log.Println("Successfully migrated")
}
