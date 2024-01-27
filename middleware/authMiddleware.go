package middlewares

import (
    // TODO: Import packages yang diperlukan
    "github.com/gin-gonic/gin"
    "github.com/ahargunyllib/task-5-pbi-btpns-Nugraha_Billy_Viandy/helpers"
    "net/http"
)

// TODO: Implementasi middleware untuk autentikasi JWT
// - Fungsi ini akan memeriksa keberadaan token JWT pada request
// - Validasi token menggunakan fungsi helper JWT
// - Jika token valid, lanjutkan ke handler selanjutnya
// - Jika tidak, kembalikan error response (misal: status 401 Unauthorized)
func JWTAuthMiddleware() gin.HandlerFunc {
    return func(context *gin.Context) {
        err := helpers.ValidateJWT(context)
        if err != nil {
            context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
            context.Abort()
            return
        }
        context.Next()
    }
}

// TODO: Tambahkan logika untuk mengekstrak dan menyimpan informasi user dari token
// - Jika token valid, ekstrak informasi user (misal: user ID) dan tambahkan ke context Gin
// - Ini memungkinkan handler selanjutnya untuk mengakses informasi user

// TODO: Pertimbangkan penanganan error yang efisien
// - Pastikan mengirimkan response yang sesuai untuk berbagai jenis error autentikasi
// - Misal: token tidak ada, token kadaluwarsa, token tidak valid

// TODO: Pertimbangkan implementasi middleware tambahan jika diperlukan
// - Misal: middleware untuk memeriksa role atau otorisasi user tertentu

