package repository

import (
	"backend_ukmik/domain"
	"backend_ukmik/model"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user model.User) error {
	result := &model.User{}
	err := r.db.Where("username = ? ", user.Username).Where("deleted_at IS NULL").First(&result).Error
	if err == nil {
		// Data duplikat ditemukan, tangani dengan sesuai
		return errors.New("username is already exist")
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		// Data tidak ditemukan, simpan data baru
		err := r.db.Create(&user).Error
		if err != nil {
			return errors.New("failed to create user")
		}
	} else {
		// Terjadi kesalahan lain saat mencari data
		return errors.New("error occurred while searching for user")
	}

	return nil
}

// cari user berdasarkan id
func (r *UserRepository) FindUserById(id int) (model.User, error) {
	result := model.User{}
	err := r.db.Table("users").Where("id = ? ", id).Where("deleted_at IS NULL").First(&result).Error
	if err != nil {
		return model.User{}, errors.New("id not found")
	}
	return result, nil
}
