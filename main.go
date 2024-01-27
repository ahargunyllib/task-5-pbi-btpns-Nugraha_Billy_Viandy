package main

import (
	// TODO: Import necessary packages
	// "github.com/gin-gonic/gin"
	// "path/to/your/router"
	// "path/to/your/database"
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

	// TODO: Setup dan inisialisasi Gin router
	// TODO: Register routes dari router/router.go
	r := router.SetupRouter()
	r.Run(":8080")
	// TODO: Jalankan server Gin
}
