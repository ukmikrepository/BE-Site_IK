package usecase

import (
	"backend_ukmik/domain"
	"backend_ukmik/model"
	"errors"
)

type CAUsecase struct {
	CARepository domain.CARepository
}

func NewCAUsecase(CARepository domain.CARepository) domain.CAUsecase {
	return &CAUsecase{
		CARepository: CARepository,
	}
}

func (c *CAUsecase) RegisterCA(clanggota model.RegCA, key int) error {
	calonAnggota := model.CA{Nama: clanggota.Nama, Email: clanggota.Email, Nim: clanggota.Nim, Jurusan: clanggota.Jurusan, Angkatan: clanggota.Angkatan, NoTlp: clanggota.NoTlp, Fakultas: clanggota.Fakultas, JKelamin: clanggota.JKelamin, Img: clanggota.Img, CreatedByUserID: uint(key)}
	jurusan := []string{"S1 - Informatika", "S1 - Sistem Informasi", "S1 - Teknik Komputer", "D3 - Sistem Informasi Akutansi", "D3 - Teknik Komputer", "D3 - Rekayasa Perangkat Lunak"}

	validateJurusan := false
	for i := range jurusan {
		if i == clanggota.Jurusan {
			validateJurusan = true
		}
	}
	if !validateJurusan {
		return errors.New("jurusan tidak ditemukan")
	}

	err := c.CARepository.DBRegisterCA(calonAnggota)
	if err != nil {
		return err
	}
	return nil
}

func (c *CAUsecase) UpdateCA(clanggota model.RegCA, idCa, key int) error {

	calonAnggota := model.CA{Nama: clanggota.Nama, Email: clanggota.Email, Nim: clanggota.Nim, Jurusan: clanggota.Jurusan, Angkatan: clanggota.Angkatan, NoTlp: clanggota.NoTlp, Fakultas: clanggota.Fakultas, JKelamin: clanggota.JKelamin, Img: clanggota.Img, UpdatedByUserID: uint(key)}
	jurusan := []string{"S1 - Informatika", "S1 - Sistem Informasi", "S1 - Teknik Komputer", "D3 - Sistem Informasi Akutansi", "D3 - Teknik Komputer", "D3 - Rekayasa Perangkat Lunak"}

	validateJurusan := false
	for i := range jurusan {
		if i == clanggota.Jurusan {
			validateJurusan = true
		}
	}
	if !validateJurusan {
		return errors.New("jurusan tidak ditemukan")
	}

	err := c.CARepository.DBUpdateCA(calonAnggota, idCa)
	if err != nil {
		return err
	}
	return nil
}

func (c *CAUsecase) ListCA(offset int, limit int) ([]model.ListCA, error) {
	dblistanggota, err := c.CARepository.DBListCA(offset, limit)
	if err != nil {
		return []model.ListCA{}, err
	}
	result := []model.ListCA{}
	for i, data := range dblistanggota {
		result = append(result, model.ListCA{No: i + 1, Img: data.Img, Nama: data.Nama, Email: data.Email, Nim: data.Nim, Jurusan: data.Jurusan, Angkatan: data.Angkatan, NoTlp: data.NoTlp})
	}
	return result, nil
}

func (c *CAUsecase) DeleteCA(idCa, key int) error {
	err := c.CARepository.DBDeleteCA(idCa, key)
	if err != nil {
		return err
	}
	return nil
}

func (c *CAUsecase) TotalCa() (int64, error) {
	total, err := c.CARepository.DBTotalCa()
	if err != nil {
		return 0, err
	}
	return total, nil
}
