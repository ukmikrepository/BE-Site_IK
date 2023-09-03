package model

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	ID       int    `gorm:"type:integer;primary_key"`
	Name     string `gorm:"type:varchar(100);not null"`
	Username string `gorm:"type:varchar(100);not null"`
	Password string `gorm:"type:varchar(100);not null"`
	Role     int    `gorm:"type:integer(10)"`

	UpdatedByUser uint `gorm:"type:integer(10);default:null"`
	DeletedByUser uint `gorm:"type:integer(10);default:null"`
}

type JSONUser struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}
