package app

// TODO: Definisikan struct User dengan field yang sesuai
// - ID (int, primary key)
// - Username (string)
// - Email (string, unik)
// - Password (string)
// - Relasi dengan model Photo (jika ada)

// TODO: Tambahkan tag untuk GORM (jika menggunakan GORM) untuk mendefinisikan struktur tabel
// Contoh: `gorm:"uniqueIndex"`

// TODO: Tambahkan validasi untuk setiap field
// - Gunakan validator, seperti Go Validator, untuk memastikan email valid, password memenuhi kriteria tertentu, dll.

// TODO: Tambahkan method-method untuk operasi yang berhubungan dengan User
// - SaveUser() untuk menyimpan user ke database
// - UpdateUser() untuk update data user
// - DeleteUser() untuk menghapus user
// - FindUserByEmail() atau FindUserByID() untuk mencari user

// TODO: Pertimbangkan implementasi fitur keamanan seperti enkripsi password
// - Gunakan bcrypt atau library serupa untuk mengenkripsi password sebelum disimpan di database

// TODO: Tambahkan fungsi yang berhubungan dengan autentikasi dan otorisasi
// - Misalnya, fungsi untuk memeriksa apakah email sudah terdaftar, fungsi untuk membandingkan password terenkripsi

// TODO: Pastikan bahwa setiap interaksi dengan database ditangani dengan penanganan error yang tepat

// TODO: Pertimbangkan relasi dengan model Photo
// - Jika aplikasi memiliki fitur foto, pastikan struktur dan method mendukung relasi antara User dan Photo
