package repository

import (
	"backend_ukmik/domain"
	"backend_ukmik/model"
	"errors"

	"gorm.io/gorm"
)

type CARepository struct {
	db *gorm.DB
}

func NewCARepository(db *gorm.DB) domain.CARepository {
	return &CARepository{
		db: db,
	}
}

func (c *CARepository) DBRegisterCA(clanggota model.CA) error {
	result := &model.CA{}
	err := c.db.Where("nim = ? ", clanggota.Nim).Where("deleted_at IS NULL").First(&result).Error
	if err == nil {
		// Data duplikat ditemukan, tangani dengan sesuai
		return errors.New("nim is already exist")
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		// Data tidak ditemukan, simpan data baru
		err := c.db.Create(&clanggota).Error
		if err != nil {
			return errors.New("failed to create calon anggota")
		}
	} else {
		// Terjadi kesalahan lain saat mencari data
		return errors.New("error occurred while searching for calon anggota")
	}

	return nil
}

func (c *CARepository) DBUpdateCA(clanggota model.CA, idCa int) error {
	result := &model.CA{}
	err := c.db.Where("nim = ? AND id != ? ", clanggota.Nim, idCa).Where("deleted_at IS NULL").First(&result).Error
	if err == nil {
		// Data duplikat ditemukan, tangani dengan sesuai
		return errors.New("nim is already exist")
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		// Data tidak ditemukan, simpan data baru
		err := c.db.Where("id = ?", idCa).Updates(&clanggota).Error
		if err != nil {
			return errors.New("failed to update calon anggota")
		}
	} else {
		// Terjadi kesalahan lain saat mencari data
		return errors.New("error occurred while searching for calon anggota")
	}

	return nil
}

func (c *CARepository) DBListCA(offset int, limit int) ([]model.ListCA, error) {
	result := []model.ListCA{}
	err := c.db.Table("cas").Select("id", "img", "nama", "email", "nim", "jurusan", "angkatan", "no_tlp").Where("deleted_at IS NULL").Order("cas.created_at ASC").Offset(offset).Limit(limit).Find(&result).Error
	if err != nil {
		return []model.ListCA{}, err
	}
	return result, nil
}

func (c *CARepository) DBDeleteCA(idCa, key int) error {
	org := model.CA{}
	err1 := c.db.Where("id = ?", idCa).Where("deleted_at IS NULL").First(&org).Error
	if errors.Is(err1, gorm.ErrRecordNotFound) {
		return errors.New("id unit not found")
	}
	err := c.db.Model(&model.CA{}).Where("id = ?", idCa).Updates(model.CA{DeletedByUserID: uint(key)}).Error
	if err != nil {
		return errors.New("failed to update organization deleted by user")
	}
	return c.db.Where("id =?", idCa).Delete(&model.CA{}).Error
}

func (c *CARepository) DBTotalCa() (int64, error) {
	var result int64
	err := c.db.Table("cas").Where("deleted_at IS NULL").Count(&result).Error
	if err != nil {
		return result, err
	}
	return result, nil
}
