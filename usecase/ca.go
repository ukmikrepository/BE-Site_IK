package usecase

import (
	"backend_ukmik/domain"
	"backend_ukmik/model"
	"errors"
	"math/rand"
	"strconv"
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

	err := c.CARepository.DBRegisterCA(calonAnggota)
	if err != nil {
		return err
	}
	return nil
}

func (c *CAUsecase) GetCAByID(id int) (model.ListCA, error) {
	data, err := c.CARepository.DBGetCAByID(id)
	if err != nil {
		return model.ListCA{}, err
	}
	return data, nil
}

func (c *CAUsecase) UpdateCA(clanggota model.RegCA, idCa, key int) error {
	calonAnggota := model.CA{Nama: clanggota.Nama, Email: clanggota.Email, Nim: clanggota.Nim, Jurusan: clanggota.Jurusan, Angkatan: clanggota.Angkatan, NoTlp: clanggota.NoTlp, Fakultas: clanggota.Fakultas, JKelamin: clanggota.JKelamin, Img: clanggota.Img, UpdatedByUserID: uint(key), StatusFee: clanggota.StatusFee}

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
	fakultas := []string{"", "Teknologi Informasi", "Bisnis Management Bisnis"}
	jurusan := []string{"", "D3 - Rekayasa Perangkat Lunak Aplikasi", "D3 - Sistem Informasi Akuntansi", "D3 - Teknologi Komputer", "S1 - Informatika", "S1 - Sistem Informasi", "S1 - Teknik Komputer", "S1 - Bisnis Digital", "S1 - Manajemen Ritel"}
	jenis_kelamin := []string{"", "Laki-laki", "Perempuan"}
	// status_pembayaran := []string{"Belum Lunas", "Lunas"}

	for i, data := range dblistanggota {
		// fakultas
		fakultasStr := ""
		idfakultas, err := strconv.Atoi(data.Fakultas)
		if err != nil {
			return nil, errors.New("convert jurusan failed")
		}
		fakultasStr = fakultas[idfakultas]

		// jurusan
		jurusanStr := ""
		idjurusan, err := strconv.Atoi(data.Jurusan)
		if err != nil {
			return nil, errors.New("convert jurusan failed")
		}
		jurusanStr = jurusan[idjurusan]

		// jenis kelamin
		jenis_kelaminStr := ""
		idjenis_kelamin, err := strconv.Atoi(data.JKelamin)
		if err != nil {
			return nil, errors.New("convert jenis kelamin failed")
		}
		jenis_kelaminStr = jenis_kelamin[idjenis_kelamin]

		// status pembayaran
		// status_pembayaranStr := ""
		// fmt.Println(data.StatusFee)
		// idstatus_pembayaran, err := strconv.Atoi(data.StatusFee)
		// if err != nil {
		// 	return nil, errors.New("convert status pembayaran failed")
		// }
		// status_pembayaranStr = status_pembayaran[idstatus_pembayaran]

		result = append(result, model.ListCA{No: i + 1, Id: data.Id, Img: data.Img, Nama: data.Nama, Email: data.Email, Nim: data.Nim, Fakultas: fakultasStr, Jurusan: jurusanStr, Angkatan: data.Angkatan, NoTlp: data.NoTlp, JKelamin: jenis_kelaminStr, StatusFee: data.StatusFee})

	}
	return result, nil
}

func (c *CAUsecase) ListAllCA() ([]model.CSVCA, error) {
	dblistanggota, err := c.CARepository.DBListAllCA()
	if err != nil {
		return []model.CSVCA{}, err
	}

	result := []model.CSVCA{}
	fakultas := []string{"", "Teknologi Informasi", "Bisnis Management Bisnis"}
	jurusan := []string{"", "D3 - Rekayasa Perangkat Lunak Aplikasi", "D3 - Sistem Informasi Akuntansi", "D3 - Teknologi Komputer", "S1 - Informatika", "S1 - Sistem Informasi", "S1 - Teknik Komputer", "S1 - Bisnis Digital", "S1 - Manajemen Ritel"}
	jenis_kelamin := []string{"", "Laki-laki", "Perempuan"}
	status_pembayaran := []string{"Belum Lunas", "Lunas"}

	for i, data := range dblistanggota {
		// fakultas
		fakultasStr := ""
		idfakultas, err := strconv.Atoi(data.Fakultas)
		if err != nil {
			return nil, errors.New("convert jurusan failed")
		}
		fakultasStr = fakultas[idfakultas]

		// jurusan
		jurusanStr := ""
		idjurusan, err := strconv.Atoi(data.Jurusan)
		if err != nil {
			return nil, errors.New("convert jurusan failed")
		}
		jurusanStr = jurusan[idjurusan]

		// jenis kelamin
		jenis_kelaminStr := ""
		idjenis_kelamin, err := strconv.Atoi(data.JKelamin)
		if err != nil {
			return nil, errors.New("convert jenis kelamin failed")
		}
		jenis_kelaminStr = jenis_kelamin[idjenis_kelamin]
		status_pembayaranStr := status_pembayaran[data.StatusFee]

		result = append(result, model.CSVCA{No: i + 1, Img: data.Img, Nama: data.Nama, Email: data.Email, Nim: data.Nim, Fakultas: fakultasStr, Jurusan: jurusanStr, Angkatan: data.Angkatan, NoTlp: data.NoTlp, JKelamin: jenis_kelaminStr, StatusFee: status_pembayaranStr})

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

func (c *CAUsecase) ValidateID(key int) error {
	err := c.CARepository.DBValidateID(key)
	if err != nil {
		return err
	}
	return nil
}

func (c *CAUsecase) GenerateID() string {
	const digits = "0123456789"
	randomNumber := make([]byte, 10)

	// Mengisi randomNumber dengan karakter acak dari digits
	for i := 0; i < 10; i++ {
		randomNumber[i] = digits[rand.Intn(len(digits))]
	}
	return string(randomNumber)
}
