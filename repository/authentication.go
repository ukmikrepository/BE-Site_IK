package repository

import (
	"backend_ukmik/domain"

	"gorm.io/gorm"
)

type AuthenticationRepository struct {
	db *gorm.DB
}

func NewAuthenticationRepository(db *gorm.DB) domain.AuthenticationRepository {
	return &AuthenticationRepository{
		db: db,
	}
}
