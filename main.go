package main

import (
	"log"

	"github.com/ahargunyllib/task-5-pbi-btpns-Nugraha_Billy_Viandy/database"
	"github.com/ahargunyllib/task-5-pbi-btpns-Nugraha_Billy_Viandy/router"

	"github.com/joho/godotenv"
)

func main() {
	// TODO: Inisialisasi database
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()
	database.Migrate()

	r := router.SetupRouter()
	r.Run(":8080")
}
