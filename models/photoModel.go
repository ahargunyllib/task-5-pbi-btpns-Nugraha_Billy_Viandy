package models

// import (
//     // TODO: Import necessary packages
//     "gorm.io/gorm"
// )

type Photo struct {
    ID       uint   `gorm:"primaryKey" json:"id"`              // ID sebagai primary key
    Title    string `gorm:"size:255" json:"title"`                // Title foto, panjang maksimal 255 karakter
    Caption  string `gorm:"size:255" json:"caption"`                // Caption foto, panjang maksimal 255 karakter
    PhotoUrl string `gorm:"size:255" json:"photoUrl"`                // URL foto, panjang maksimal 255 karakter
    UserID   uint   `gorm:"not null" json:"userId"`                // UserID, foreign key yang merujuk ke User
    User     User   `gorm:"foreignKey:UserID"`
}

// TODO: Implement any necessary GORM hooks or methods
// Example: BeforeSave, AfterCreate, etc.

// TODO: Add any model-level methods required
// Example: methods to handle business logic related to photos

// TODO: If any relationship with other models (like User), define them here
// Example: Photo belongs to User

// TODO: Consider implementing validations for photo fields if required
