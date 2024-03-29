package models

import (
	"html"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/asaskevich/govalidator"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"not null;unique" json:"username" form:"username" valid:"required~Username is required"`
	Email     string    `gorm:"not null;unique" json:"email" form:"email" valid:"required~Email is required,email~Invalid Email"`
	Password  string    `gorm:"not null" form:"password" valid:"required~Password is required,minstringlength(6)~Password minimum is 6 characters"`
	Photos    []Photo   `gorm:"constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	return user.validateStruct()
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	return user.validateStruct()
}

func (user *User) validateStruct() error {
	_, err := govalidator.ValidateStruct(user)
	return err
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}