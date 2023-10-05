package usecase

import (
	"backend_ukmik/domain"
	"backend_ukmik/model"
)

type DashboardUsecase struct {
	DashboardRepository domain.DashboardRepository
}

func NewDashboardUsecase(DashboardRepository domain.DashboardRepository) domain.DashboardUsecase {
	return &DashboardUsecase{
		DashboardRepository: DashboardRepository,
	}
}

func (u *DashboardUsecase) Dashboarad(key int) (model.User, error) {
	return model.User{}, nil
}
