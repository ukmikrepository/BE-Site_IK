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

	/*
		IDENTIFIED ROLE
		0 = Anggota
		1 = SuperAdmin
		2 = Admin
		3 = Pengurus Harian
		4 = Sekretaris
		5 = Bendahara
		6 = Keanggotaan
		7 = Pendidikan
		8 = Minat Bakat
		9 = Publikasi
		10 = Humas

	*/
	StrRole := []string{"Anggota", "Super Admin", "Admin", "Pengurus Harian", "Sekretaris", "Bendahara", "Keanggotaan", "Pendidikan", "Minat Bakat", "Publikasi", "Humas"}

	// Generate Token
	token, _ := utils.GenerateToken(config.TokenExpiresIn, dbUser.ID, StrRole[dbUser.Role], config.TokenSecret)
	return token, nil
}
