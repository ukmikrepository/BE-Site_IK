package usecase

import (
	"backend_ukmik/domain"
	"backend_ukmik/model"
	"fmt"
	"math/rand"
)

type UserUsecase struct {
	UserRepository domain.UserRepository
}

func NewUserUsecase(UserRepository domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{
		UserRepository: UserRepository,
	}
}

func (u *UserUsecase) CreateUser(user model.JSONUser, key int) error {
	dbuser := model.User{Name: user.Name, Username: user.Username, Password: user.Password, NRA: user.NRA, Image: user.Img, Role: user.Role}
	err := u.UserRepository.CreateUser(dbuser)
	if err != nil {
		return err
	}

	fmt.Println(dbuser)
	return nil
}

func (u *UserUsecase) GenerateID() string {
	const digits = "0123456789"
	randomNumber := make([]byte, 10)

	// Mengisi randomNumber dengan karakter acak dari digits
	for i := 0; i < 10; i++ {
		randomNumber[i] = digits[rand.Intn(len(digits))]
	}
	return string(randomNumber)
}
