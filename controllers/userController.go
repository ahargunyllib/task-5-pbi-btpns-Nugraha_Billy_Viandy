package controllers

import (
	// TODO: Import necessary packages
	"net/http"

	"github.com/ahargunyllib/task-5-pbi-btpns-Nugraha_Billy_Viandy/app"
	"github.com/ahargunyllib/task-5-pbi-btpns-Nugraha_Billy_Viandy/database"
	"github.com/ahargunyllib/task-5-pbi-btpns-Nugraha_Billy_Viandy/models"
    "github.com/ahargunyllib/task-5-pbi-btpns-Nugraha_Billy_Viandy/helpers"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(context *gin.Context) {
	var input app.RegisterInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	savedUser := database.Database.Create(&user)

	if savedUser.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": savedUser.Error.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message":"User berhasil dibuat"})
}

// TODO: Implementasi fungsi untuk login user (POST /users/login)
// - Terima email dan password dari request
// - Validasi data
// - Verifikasi user dan password
// - Buat dan return JWT token jika login sukses
// - Handle error dan return response yang sesuai
func Login(context *gin.Context) {
    var input app.LoginInput

    if err := context.ShouldBindJSON(&input); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    err := database.Database.Where("username=?", input.Username).Find(&user).Error
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    jwt, err := helpers.GenerateJWT(user)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusOK, gin.H{"jwt": jwt})
}

// TODO: Implementasi fungsi untuk update user (PUT /users/:userId)
// - Validasi JWT
// - Terima dan validasi data update dari request
// - Update data user di database
// - Handle error dan return response yang sesuai

// TODO: Implementasi fungsi untuk menghapus user (DELETE /users/:userId)
// - Validasi JWT
// - Verifikasi bahwa user yang login adalah yang ingin dihapus
// - Hapus user dari database
// - Handle error dan return response yang sesuai
