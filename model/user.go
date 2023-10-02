package model

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	ID       int    `gorm:"type:integer;primary_key"`
	Name     string `gorm:"type:varchar(100);not null"`
	Username string `gorm:"type:varchar(100);not null"`
	Password string `gorm:"type:varchar(100);not null"`
	Role     int    `gorm:"type:integer"`
	NRA      string `gorm:"type:varchar(100);not null"`
	Image    string

	CreateByyUser uint `gorm:"type:integer;default:null"`
	UpdatedByUser uint `gorm:"type:integer;default:null"`
	DeletedByUser uint `gorm:"type:integer;default:null"`
}

type Member struct {
	ID     int `gorm:"type:integer;primary_key"`
	UserID int // Ini adalah field untuk menyimpan ID user yang terkait dengan member ini
	// Tambahkan field lain yang relevan untuk Member di sini
}

type JSONUser struct {
	Name     string `form:"name"`
	Username string `form:"username"`
	Password string `form:"password"`
	Role     int    `form:"role"`
	NRA      string `form:"nra"`
	Img      string `form:"img"`
}

type ResLogin struct {
	Res   Response
	Token string `json:"token"`
}
