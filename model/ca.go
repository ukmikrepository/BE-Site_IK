package model

import (
	"gorm.io/gorm"
)

type CA struct {
	*gorm.Model
	ID        int    `gorm:"primaryKey"`
	Nama      string `gorm:"type:varchar(50)"`
	Email     string `gorm:"type:varchar(50)"`
	Img       string
	Nim       string `gorm:"type:varchar(11)"`
	Fakultas  int
	Jurusan   int
	Angkatan  string `gorm:"type:varchar(5)"`
	NoTlp     string `gorm:"type:varchar(15)"`
	JKelamin  int
	StatusFee int

	CreatedByUserID uint
	CreatedByUser   User `gorm:"foreignKey:CreatedByUserID"`
	UpdatedByUserID uint `gorm:"default:null"`
	UpdatedByUser   User `gorm:"foreignKey:UpdatedByUserID"`
	DeletedByUserID uint `gorm:"default:null"`
	DeletedByUser   User `gorm:"foreignKey:DeletedByUserID;constraint:OnDelete:CASCADE"`
}

type RegCA struct {
	Nama      string `form:"nama"`
	Email     string `form:"email"`
	Nim       string `form:"nim"`
	Fakultas  int    `form:"fakultas"`
	Jurusan   int    `form:"jurusan"`
	Angkatan  string `form:"angkatan"`
	NoTlp     string `form:"no_telp"`
	JKelamin  int    `form:"j_kelamin"`
	Img       string `form:"img"`
	StatusFee int    `form:"status_fee"`
}

type ListCA struct {
	No        int    `json:"no"`
	Id        int    `json:"id"`
	Img       string `json:"image"`
	Nama      string `json:"nama"`
	Email     string `json:"email"`
	Nim       string `json:"nim"`
	Fakultas  string `json:"fakultas"`
	Jurusan   string `json:"jurusan"`
	Angkatan  string `json:"angkatan"`
	NoTlp     string `json:"no_telpon"`
	JKelamin  string `json:"jenis_kelamin"`
	StatusFee int    `json:"status_fee"`
}

type CSVCA struct {
	No        int    `json:"no"`
	Img       string `json:"image"`
	Nama      string `json:"nama"`
	Email     string `json:"email"`
	Nim       string `json:"nim"`
	Fakultas  string `json:"fakultas"`
	Jurusan   string `json:"jurusan"`
	Angkatan  string `json:"angkatan"`
	NoTlp     string `json:"no_telpon"`
	JKelamin  string `json:"jenis_kelamin"`
	StatusFee string `json:"status_fee"`
}

type ResponseListCA struct {
	Response Response
	Meta     Meta
	Data     []ListCA
}
