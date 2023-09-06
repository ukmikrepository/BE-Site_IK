package domain

import "backend_ukmik/model"

type AuthenticationRepository interface {
	DBValidateUser(user model.Login) (model.User, error)
}

type AuthenticationUsecase interface {
	ValidateUser(user model.Login) (string, error)
}
