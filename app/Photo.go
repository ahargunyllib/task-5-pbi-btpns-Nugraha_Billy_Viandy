package app

type CreatePhoto struct {
    Title    string `gorm:"size:255" json:"title"`
    Caption  string `gorm:"size:255" json:"caption"`
    PhotoUrl string `gorm:"size:255" json:"photoUrl"`
}
