package domain

import "backend_ukmik/model"

type UserRepository interface {
	CreateUser(user model.User) error
}

type UserUsecase interface {
	CreateUser(user model.JSONUser) error
}
