package repository

import (
	"backend_ukmik/domain"
	"backend_ukmik/model"
	"errors"
	"fmt"
	"os"
	"path/filepath"

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

func (c *CARepository) DBGetCAByID(id int) (model.ListCA, error) {
	result := model.ListCA{}
	err := c.db.Table("cas").Select("id", "img", "nama", "email", "nim", "fakultas", "jurusan", "angkatan", "no_tlp", "j_kelamin", "status_fee").Where("deleted_at IS NULL").Order("cas.created_at ASC").Find(&result).Error
	if err != nil {
		return model.ListCA{}, err
	}
	return result, nil
}

func (c *CARepository) DBUpdateCA(clanggota model.CA, idCa int) error {
	result := &model.CA{}
	err := c.db.Where("id != ? ", idCa).Where("deleted_at IS NULL").First(&result).Error
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
	err := c.db.Table("cas").Select("id", "img", "nama", "email", "nim", "fakultas", "jurusan", "angkatan", "no_tlp", "j_kelamin", "status_fee").Where("deleted_at IS NULL").Order("cas.created_at ASC").Offset(offset).Limit(limit).Find(&result).Error
	if err != nil {
		return []model.ListCA{}, err
	}
	return result, nil
}

func (c *CARepository) DBListAllCA() ([]model.ListCA, error) {
	result := []model.ListCA{}
	err := c.db.Table("cas").Select("id", "img", "nama", "email", "nim", "fakultas", "jurusan", "angkatan", "no_tlp", "j_kelamin", "status_fee").Where("deleted_at IS NULL").Order("cas.created_at ASC").Find(&result).Error
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
	fmt.Println(org)
	backupImg := "backup-" + org.Img
	err := c.db.Model(&model.CA{}).Where("id = ?", idCa).Updates(model.CA{DeletedByUserID: uint(key), Img: backupImg}).Error
	if err != nil {
		return errors.New("failed to update ca deleted by user")
	}

	destination := "uploads/image/ca/2023/"

	oldFilePath := filepath.Join(destination, org.Img)
	newFilePath := filepath.Join(destination, backupImg)

	err = os.Rename(oldFilePath, newFilePath)
	if err != nil {
		return err
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

func (c *CARepository) DBValidateID(key int) error {
	var count int64
	if err := c.db.Table("cas").Where("id = ?", key).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("ID tidak ditemukan") // Menghasilkan error jika ID tidak ditemukan
	}
	return nil
}
