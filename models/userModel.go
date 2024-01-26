package models

import (
    // TODO: Import necessary packages
    "time"
    "gorm.io/gorm"
    "github.com/asaskevich/govalidator"
)

// User struct represents the user model
type User struct {
        ID          uint        `gorm:"primaryKey" json:"id"`
        Username    string      `gorm:"not null;unique" json:"username" form:"username" valid:"required~Username is required"`
        Email       string      `gorm:"not null;unique" json:"email" form:"email" valid:"required~Email is required,email~Invalid Email"`
        Password    string      `gorm:"size:6;not null" form:"password" valid:"required~Password is required,minstringlength(6)~Password minimum is 6 characters"`
        Photos      []Photo     `gorm:"constraint:OnDelete:CASCADE;"`
        CreatedAt   time.Time   `gorm:"autoCreateTime" json:"created_at"`
        UpdatedAt   time.Time   `gorm:"autoUpdateTime" json:"updated_at"` 
}

// BeforeCreate digunakan untuk validasi sebelum data User dibuat.
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	return user.validateStruct()
}

// BeforeUpdate digunakan untuk validasi sebelum data User diperbarui.
func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	return user.validateStruct()
}

// validateStruct digunakan untuk melakukan validasi struct menggunakan govalidator.
func (user *User) validateStruct() error {
	_, err := govalidator.ValidateStruct(user)
	return err
}
