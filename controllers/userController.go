package controllers

import (
	// TODO: Import necessary packages
	"net/http"
	"strconv"

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

	var user models.User
	
	if err := database.Database.Where("email = ?", input.Email).First(&user).Error; err == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Email sudah terdaftar"})
		context.Abort()
		return
	}

	// cek apakah username sudah terdaftar
	if err := database.Database.Where("username = ?", input.Username).First(&user).Error; err == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Username sudah terdaftar"})
		context.Abort()
		return
	}

	user = models.User{
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

func Login(context *gin.Context) {
    var input app.LoginInput

    if err := context.ShouldBindJSON(&input); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    err := database.Database.Where("email=?", input.Email).Find(&user).Error
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Email salah"})
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Password salah"})
        return
    }

    jwt, err := helpers.GenerateJWT(user)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Tidak dapat generete token"})
        return
    }

    context.JSON(http.StatusOK, gin.H{"jwt": jwt})
}

func Update(context *gin.Context) {
	userId, err := strconv.Atoi(context.Param("userId"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "ID pengguna tidak valid"})
		return
	}

	var input app.UpdateInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newUser models.User

	if err := database.Database.Where("email = ? AND id != ?", input.Email, userId).First(&newUser).Error; err == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Email sudah terdaftar"})
		context.Abort()
		return
	}

	if err := database.Database.Where("username = ? AND id != ?", input.Username, userId).First(&newUser).Error; err == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Username sudah terdaftar"})
		context.Abort()
		return
	}

	user, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userId != int(user.ID) {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Tidak diizinkan"})
		context.Abort()
		return
	}

	if err := database.Database.Where("id = ?", user.ID).First(&newUser).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Pengguna tidak ditemukan"})
		return
	}

	if err := database.Database.First(&newUser, userId).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Pengguna tidak ditemukan"})
		return
	}

	newUser.Email = input.Email
	newUser.Username = input.Username
	newUser.Password = input.Password

	if err := database.Database.Save(&newUser).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Berhasil mengupdate pengguna"})
}

func Delete(context *gin.Context){
	var user models.User
	userId, err := strconv.Atoi(context.Param("userId"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "ID pengguna tidak valid"})
		return
	}

	currentUser, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userId != int(currentUser.ID) {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Tidak diizinkan"})
		context.Abort()
		return
	}

	if err := database.Database.Where("id = ?", currentUser.ID).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Pengguna tidak ditemukan"})
		return
	}

	if err := database.Database.First(&user, userId).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Pengguna tidak ditemukan"})
		return
	}	

	if err := database.Database.Delete(&user).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Berhasil menghapus pengguna"})
}