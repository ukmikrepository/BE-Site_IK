package usecase

import (
	"backend_ukmik/config"
	"backend_ukmik/domain"
	"backend_ukmik/model"
	"backend_ukmik/utils"
)

type AuthenticationUsecase struct {
	AuthenticationRepository domain.AuthenticationRepository
}

func NewAuthenticationUsecase(AuthenticationRepository domain.AuthenticationRepository) domain.AuthenticationUsecase {
	return &AuthenticationUsecase{
		AuthenticationRepository: AuthenticationRepository,
	}
}

func (a *AuthenticationUsecase) ValidateUser(user model.Login) (string, error) {
	dbUser, err := a.AuthenticationRepository.DBValidateUser(user)
	if err != nil {
		return "", err
	}

	config, _ := config.LoadConfigEnv(".")

	// Generate Token
	token, _ := utils.GenerateToken(config.TokenExpiresIn, dbUser.ID, config.TokenSecret)
	return token, nil
}
