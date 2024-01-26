package app

// TODO: Definisikan struct Photo dengan field yang sesuai
// - ID (int, primary key)
// - Title (string)
// - Caption (string)
// - PhotoUrl (string)
// - UserID (int, foreign key yang merujuk ke User)

// TODO: Tambahkan tag untuk GORM (jika menggunakan GORM) untuk mendefinisikan struktur tabel
// Contoh: `gorm:"foreignKey:UserID"`

// TODO: Tambahkan validasi untuk setiap field (gunakan validator, seperti Go Validator jika perlu)

// TODO: Jika perlu, tambahkan method-method khusus yang berhubungan dengan Photo
// Contoh: 
// - SavePhoto() untuk menyimpan photo ke database
// - UpdatePhoto() untuk update data photo
// - DeletePhoto() untuk menghapus photo
// - FindPhotoByID() untuk mencari photo berdasarkan ID

// TODO: Tambahkan fungsi untuk menghubungkan Photo dengan User (jika diperlukan)
// - Misalnya, fungsi untuk mendapatkan semua foto dari user tertentu

// TODO: Pertimbangkan menambahkan fungsi untuk handling upload file
// - Jika aplikasi Anda memungkinkan pengguna meng-upload foto, pertimbangkan implementasi di sini

// TODO: Pastikan bahwa setiap interaksi dengan database ditangani dengan penanganan error yang tepat
