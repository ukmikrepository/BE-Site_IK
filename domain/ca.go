package domain

import "backend_ukmik/model"

type CARepository interface {
	DBRegisterCA(clanggota model.CA) error
	DBUpdateCA(clanggota model.CA, IdCa int) error
	DBListCA(offset int, limit int) ([]model.ListCA, error)
	DBTotalCa() (int64, error)
	DBDeleteCA(idCa, key int) error
	DBValidateID(int) error
}

type CAUsecase interface {
	RegisterCA(clanggota model.RegCA, key int) error
	UpdateCA(clanggota model.RegCA, IdCa, key int) error
	ListCA(offset int, limit int) ([]model.ListCA, error)
	TotalCa() (int64, error)
	DeleteCA(idCa, key int) error
	ValidateID(int) error
}
