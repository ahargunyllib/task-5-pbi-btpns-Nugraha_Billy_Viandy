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

	publicRoutes := router.Group("/users")
	publicRoutes.POST("/register", controllers.Register)
	publicRoutes.POST("/login", controllers.Login)
	publicRoutes.PATCH("/:userId", controllers.Update)
	publicRoutes.DELETE("/:userId", controllers.Delete)

	protectedRoutes := router.Group("/photos")
	protectedRoutes.Use(middlewares.JWTAuthMiddleware())
	protectedRoutes.POST("/", controllers.AddPhoto)
	protectedRoutes.GET("/", controllers.GetAllPhotos)
	protectedRoutes.PATCH("/:photoId", controllers.UpdatePhoto)
	protectedRoutes.DELETE("/:photoId", controllers.DeletePhoto)

	// Return router
	return router
}
