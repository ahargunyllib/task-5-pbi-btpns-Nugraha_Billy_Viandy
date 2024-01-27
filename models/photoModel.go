package models

// import (
//     // TODO: Import necessary packages
//     "gorm.io/gorm"
// )

type Photo struct {
    ID       uint   `gorm:"primaryKey" json:"id"`              
    Title    string `gorm:"size:255" json:"title"`                
    Caption  string `gorm:"size:255" json:"caption"`                
    PhotoUrl string `gorm:"size:255" json:"photoUrl"`                
    UserID   uint   `gorm:"not null" json:"userId"`                
    User     User   `gorm:"foreignKey:UserID"`
}