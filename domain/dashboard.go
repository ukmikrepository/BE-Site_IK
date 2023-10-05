package domain

import "backend_ukmik/model"

type DashboardRepository interface {
	DBDashboarad(int) (model.User, error)
}

type DashboardUsecase interface {
	Dashboarad(int) (model.User, error)
}
