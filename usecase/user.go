package usecase

import (
	"backend_ukmik/domain"
	"backend_ukmik/model"
)

type UserUsecase struct {
	UserRepository domain.UserRepository
}

func NewUserUsecase(UserRepository domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{
		UserRepository: UserRepository,
	}
}

func (u *UserUsecase) CreateUser(user model.JSONUser) error {
	dbuser := model.User{Name: user.Name, Username: user.Username, Password: user.Password, Role: user.Role}
	err := u.UserRepository.CreateUser(dbuser)
	if err != nil {
		return err
	}
	return nil
}
