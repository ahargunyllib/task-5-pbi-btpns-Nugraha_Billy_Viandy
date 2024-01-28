package controllers

import (
	"net/http"
	"strconv"

	"github.com/ahargunyllib/task-5-pbi-btpns-Nugraha_Billy_Viandy/app"
	"github.com/ahargunyllib/task-5-pbi-btpns-Nugraha_Billy_Viandy/database"
	"github.com/ahargunyllib/task-5-pbi-btpns-Nugraha_Billy_Viandy/helpers"
	"github.com/ahargunyllib/task-5-pbi-btpns-Nugraha_Billy_Viandy/models"
	"github.com/gin-gonic/gin"
)

func AddPhoto(context *gin.Context) {
	var input app.CreatePhoto
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    photo := models.Photo {
        UserID:     user.ID,
        PhotoUrl:   input.PhotoUrl,
        Title:      input.Title,
        Caption:    input.Caption,
		User: 		user,
    }

	savedPhoto := database.Database.Create(&photo)

	if savedPhoto.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": savedPhoto.Error.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message":"Photo berhasil dibuat"})
}

func GetAllPhotos(context *gin.Context) {
	user, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var photos models.Photo
	err = database.Database.Table("photos").Select("photos.id, photos.title, photos.caption, photos.photo_url, photos.user_id, user").Where("photos.user_id = ?", user.ID).Scan(&photos).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": photos})
}

func UpdatePhoto(context *gin.Context) {
	photoId, err := strconv.Atoi(context.Param("photoId"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "ID foto tidak valid"})
		return
	}

	var input app.UpdatePhoto
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var photo models.Photo

	if err := database.Database.Where("id = ? AND user_id = ?", photoId, user.ID).First(&photo).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Foto tidak ditemukan"})
		return
	}

	photo.Title = input.Title
	photo.Caption = input.Caption
	photo.PhotoUrl = input.PhotoUrl
	if err := database.Database.Save(&photo).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Foto berhasil diupdate"})
}

// TODO: Implementasi fungsi untuk menghapus photo (DELETE /photos/:photoId)
// - Validasi JWT dan ownership foto
// - Hapus photo dari database
// - Handle error dan return response yang sesuai
func DeletePhoto(context *gin.Context) {
	var photo models.Photo

	photoId, err := strconv.Atoi(context.Param("photoId"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "ID foto tidak valid"})
		return
	}

	user, err := helpers.CurrentUser(context)	
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	if err := database.Database.First(&photo, photoId).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Foto tidak ditemukan"})
		return
	}

	if photo.UserID != user.ID {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Anda tidak memiliki akses untuk menghapus foto ini"})
		return
	}

	if err := database.Database.Delete(&photo).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Foto berhasil dihapus"})
}