package database

import (
    "fmt"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var Database *gorm.DB

// TODO: Buat fungsi untuk menginisialisasi koneksi ke database
// - Konfigurasikan koneksi database (misalnya, baca konfigurasi dari file atau environment variables)
// - Buka koneksi ke database menggunakan GORM
// - Handle error koneksi
func Connect() {
    var err error
    host := os.Getenv("DB_HOST")
    username := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    databaseName := os.Getenv("DB_NAME")
    port := os.Getenv("DB_PORT")

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, username, password, databaseName, port)
    Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        panic(err)
    } else {
        fmt.Println("Successfully connected to the database")
    }
}

// TODO: Implementasi fungsi migrasi
// - Gunakan GORM untuk migrasi skema database
// - Definisikan struktur tabel berdasarkan models (User, Photo, dll.)
// - Pastikan untuk menangani error dan kasus ketika tabel sudah ada

// TODO: Jika perlu, tambahkan fungsi untuk seed data awal ke database
// - Gunakan untuk pengujian atau setup awal database dengan data sampel

// TODO: Pertimbangkan pembuatan fungsi untuk menutup koneksi database
// - Gunakan untuk gracefully shutdown aplikasi Anda

// TODO: Jika aplikasi Anda memerlukan, pertimbangkan implementasi fitur seperti connection pooling, transaction handling, dll.

