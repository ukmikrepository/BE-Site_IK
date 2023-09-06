package repository

import (
	"backend_ukmik/domain"
	"backend_ukmik/model"
	"fmt"

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

func (a *AuthenticationRepository) DBValidateUser(user model.Login) (model.User, error) {
	result := model.User{}
	err := a.db.Table("users").Where("username = ? AND password = ?", user.Username, user.Password).First(&result).Error
	if err != nil {
		return model.User{}, fmt.Errorf("username and password are not valid")
	}
	// fmt.Println("DB result:", result)
	return result, nil
}
