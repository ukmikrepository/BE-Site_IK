package repository

import (
	"backend_ukmik/domain"
	"backend_ukmik/model"

	"gorm.io/gorm"
)

type Dashboardepository struct {
	db *gorm.DB
}

func NewDashboardepository(db *gorm.DB) domain.DashboardRepository {
	return &Dashboardepository{
		db: db,
	}
}

func (r *Dashboardepository) DBDashboarad(key int) (model.User, error) {
	return model.User{}, nil
}
