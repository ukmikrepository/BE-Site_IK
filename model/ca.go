package model

import (
	"gorm.io/gorm"
)

type CA struct {
	*gorm.Model
	ID       int    `gorm:"primaryKey"`
	Nama     string `gorm:"type:varchar(50)"`
	Email    string `gorm:"type:varchar(50)"`
	Img      string
	Nim      string `gorm:"type:varchar(11)"`
	Fakultas int
	Jurusan  int
	Angkatan string `gorm:"type:varchar(5)"`
	NoTlp    string `gorm:"type:varchar(15)"`
	JKelamin int

	CreatedByUserID uint
	CreatedByUser   User `gorm:"foreignKey:CreatedByUserID"`
	UpdatedByUserID uint `gorm:"default:null"`
	UpdatedByUser   User `gorm:"foreignKey:UpdatedByUserID"`
	DeletedByUserID uint `gorm:"default:null"`
	DeletedByUser   User `gorm:"foreignKey:DeletedByUserID;constraint:OnDelete:CASCADE"`
}

type RegCA struct {
	Nama     string `form:"nama" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Nim      string `form:"nim" binding:"required"`
	Fakultas int    `form:"fakultas" binding:"required"`
	Jurusan  int    `form:"jurusan" binding:"required"`
	Angkatan string `form:"angkatan" binding:"required"`
	NoTlp    string `form:"no_telp" binding:"required"`
	JKelamin int    `form:"j_kelamin" binding:"required"`
	Img      string `form:"img"`
}

type ListCA struct {
	No       int    `json:"no"`
	Id       int    `json:"id"`
	Img      string `json:"image"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Nim      string `json:"nim"`
	Fakultas string `json:"fakultas"`
	Jurusan  string `json:"jurusan"`
	Angkatan string `json:"angkatan"`
	NoTlp    string `json:"no_telpon"`
	JKelamin string `json:"jenis_kelamin"`
}

type ResponseListCA struct {
	Response Response
	Meta     Meta
	Data     []ListCA
}
