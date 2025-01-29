package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Phone    string `json:"phone" gorm:"unique"`
	Verified bool   `json:"verified" gorm:"default:false"`
}
