package router

import (
	// TODO: Import packages yang diperlukan
	"github.com/gin-gonic/gin"

	"github.com/ahargunyllib/task-5-pbi-btpns-Nugraha_Billy_Viandy/controllers"
	// "github.com/ahargunyllib/task-5-pbi-btpns-Nugraha_Billy_Viandy/middleware"
)

// TODO: Setup function untuk konfigurasi dan inisialisasi router Gin
func SetupRouter() *gin.Engine {
	// Inisialisasi router Gin
	router := gin.Default()

	// TODO: Konfigurasikan middleware global jika ada
	// Contoh: r.Use(middlewares.AuthMiddleware())
	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controllers.Register)
	// publicRoutes.POST("/login", controllers.Login)
	// TODO: Definisikan rute untuk User
	// Contoh:
	// userGroup := r.Group("/users")
	// {
	//     userGroup.POST("/register", controllers.RegisterUser)
	//     userGroup.POST("/login", controllers.LoginUser)
	//     userGroup.PUT("/:userId", middlewares.AuthMiddleware(), controllers.UpdateUser)
	//     userGroup.DELETE("/:userId", middlewares.AuthMiddleware(), controllers.DeleteUser)
	// }

	// TODO: Definisikan rute untuk Photo
	// Contoh:
	// photoGroup := r.Group("/photos")
	// photoGroup.Use(middlewares.AuthMiddleware()) // Middleware bisa diaplikasikan pada grup
	// {
	//     photoGroup.POST("/", controllers.AddPhoto)
	//     photoGroup.GET("/", controllers.GetAllPhotos)
	//     photoGroup.PUT("/:photoId", controllers.UpdatePhoto)
	//     photoGroup.DELETE("/:photoId", controllers.DeletePhoto)
	// }

	// TODO: Tambahkan rute lain jika diperlukan

	// Return router
	return router
}
