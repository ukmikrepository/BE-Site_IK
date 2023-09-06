package model

import "gorm.io/gorm"

type CA struct {
	*gorm.Model
	ID       int `gorm:"primaryKey"`
	Nama     string
	Email    string
	Img      string
	Nim      string
	Jurusan  int
	Angkatan string
	NoTlp    string

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
	Jurusan  int    `form:"jurusan" binding:"required"`
	Angkatan string `form:"angkatan" binding:"required"`
	NoTlp    string `form:"no_telp" binding:"required"`
}

type ListCA struct {
	No       int
	Img      string
	Nama     string
	Email    string
	Nim      string
	Jurusan  int
	Angkatan string
	NoTlp    string
}

type ResponseListCA struct {
	Response Response
	Meta     Meta
	Data     []ListCA
}
