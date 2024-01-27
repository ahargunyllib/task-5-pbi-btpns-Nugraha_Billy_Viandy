package helpers

import (
	// TODO: Import package JWT
	"os"
    "strconv"
    "time"

	"github.com/golang-jwt/jwt"
    "github.com/ahargunyllib/task-5-pbi-btpns-Nugraha_Billy_Viandy/models"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

// TODO: Implementasi fungsi untuk membuat JWT token
// - Fungsi ini akan menerima data pengguna atau claims dan mengembalikan token JWT
// - Tentukan durasi masa berlaku token
func GenerateJWT(user models.User) (string, error) {
    tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id":  user.ID,
        "iat": time.Now().Unix(),
        "eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
    })
    return token.SignedString(privateKey)
}

// TODO: Implementasi fungsi untuk memvalidasi dan memparse JWT token
// - Fungsi ini akan memverifikasi token yang diberikan dan memparse claims
// - Pastikan untuk menangani kasus token yang kadaluwarsa atau tidak valid

// TODO: Pertimbangkan keamanan dan best practices dalam pembuatan token
// - Misalnya, gunakan signing method yang aman dan simpan secret key dengan baik

// TODO: Pertimbangkan penanganan error yang efektif
// - Pastikan bahwa fungsi mengembalikan error yang jelas dan bermanfaat jika terjadi masalah
