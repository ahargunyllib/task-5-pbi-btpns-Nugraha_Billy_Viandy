package models

import (
    // TODO: Import necessary packages
    "gorm.io/gorm"
)

// Photo struct represents the photo model
type Photo struct {
    // TODO: Define Photo model fields
    // Example: ID, Title, Caption, PhotoUrl, UserID
    // Include GORM tags for database properties
    gorm.Model
    ID       uint   `gorm:"primaryKey"`              // ID sebagai primary key
    Title    string `gorm:"size:255"`                // Title foto, panjang maksimal 255 karakter
    Caption  string `gorm:"size:255"`                // Caption foto, panjang maksimal 255 karakter
    PhotoUrl string `gorm:"size:255"`                // URL foto, panjang maksimal 255 karakter
    UserID   uint   `gorm:"not null"`                // UserID, foreign key yang merujuk ke User
    User     User   `gorm:"foreignKey:UserID"`
}

// TODO: Implement any necessary GORM hooks or methods
// Example: BeforeSave, AfterCreate, etc.

// TODO: Add any model-level methods required
// Example: methods to handle business logic related to photos

// TODO: If any relationship with other models (like User), define them here
// Example: Photo belongs to User

// TODO: Consider implementing validations for photo fields if required
