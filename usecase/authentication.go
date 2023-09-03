package usecase

import (
	"backend_ukmik/domain"
)

type AuthenticationUsecase struct {
	AuthenticationRepository domain.AuthenticationRepository
}

func NewAuthenticationUsecase(AuthenticationRepository domain.AuthenticationRepository) domain.AuthenticationUsecase {
	return &AuthenticationUsecase{
		AuthenticationRepository: AuthenticationRepository,
	}
}
