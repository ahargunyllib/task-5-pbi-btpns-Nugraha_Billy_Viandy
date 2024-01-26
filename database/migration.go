package database

import (
	"log"

	"github.com/ahargunyllib/task-5-pbi-btpns-Nugraha_Billy_Viandy/models"
)

func Migrate() {
    Database.AutoMigrate(&models.User{})
    log.Println("Database Migration Completed!")
    // TODO: Handle any migration errors
}

