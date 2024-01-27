package router

import (
	// TODO: Import packages yang diperlukan
	"github.com/gin-gonic/gin"

	"github.com/ahargunyllib/task-5-pbi-btpns-Nugraha_Billy_Viandy/controllers"
	middlewares "github.com/ahargunyllib/task-5-pbi-btpns-Nugraha_Billy_Viandy/middleware"
)

// TODO: Setup function untuk konfigurasi dan inisialisasi router Gin
func SetupRouter() *gin.Engine {
	// Inisialisasi router Gin
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controllers.Register)
	publicRoutes.POST("/login", controllers.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middlewares.JWTAuthMiddleware())
	protectedRoutes.POST("/photos", controllers.AddPhoto)
	protectedRoutes.GET("/photos", controllers.GetAllPhotos)

	// Return router
	return router
}
