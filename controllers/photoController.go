package controllers

import (
	"net/http"

	"github.com/ahargunyllib/task-5-pbi-btpns-Nugraha_Billy_Viandy/app"
	"github.com/ahargunyllib/task-5-pbi-btpns-Nugraha_Billy_Viandy/database"
	"github.com/ahargunyllib/task-5-pbi-btpns-Nugraha_Billy_Viandy/helpers"
	"github.com/ahargunyllib/task-5-pbi-btpns-Nugraha_Billy_Viandy/models"
	"github.com/gin-gonic/gin"
	// TODO: Import necessary packages
)

// TODO: Implementasi fungsi untuk menambahkan photo baru (POST /photos)
// - Validasi JWT
// - Terima data photo dari request
// - Validasi data
// - Simpan photo ke database dengan relasi ke user
// - Handle error dan return response yang sesuai
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
    }

	savedPhoto := database.Database.Create(&photo)

	if savedPhoto.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": savedPhoto.Error.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message":"Photo berhasil dibuat"})
}

// TODO: Implementasi fungsi untuk mendapatkan semua photo (GET /photos)
// - Fetch data photo dari database
// - Return data photo
// - Handle error dan return response yang sesuai
func GetAllPhotos(context *gin.Context) {
	user, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user.Photos})
}

// TODO: Implementasi fungsi untuk update photo (PUT /photos/:photoId)
// - Validasi JWT dan ownership foto
// - Terima dan validasi data update dari request
// - Update photo di database
// - Handle error dan return response yang sesuai

// TODO: Implementasi fungsi untuk menghapus photo (DELETE /photos/:photoId)
// - Validasi JWT dan ownership foto
// - Hapus photo dari database
// - Handle error dan return response yang sesuai
