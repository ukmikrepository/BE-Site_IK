package controller

import (
	"backend_ukmik/domain"
	"backend_ukmik/model"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CAController struct {
	CAUsecase domain.CAUsecase
}

func NewCAController(CAUsecase domain.CAUsecase) *CAController {
	return &CAController{CAUsecase}
}

func (ca *CAController) RegisterCA(c *gin.Context) {
	var clanggota model.RegCA

	if err := c.ShouldBind(&clanggota); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "field is invalid",
		})
		return
	}

	if len(clanggota.Nama) >= 50 || len(clanggota.Email) >= 50 || len(clanggota.Nim) >= 11 || clanggota.Jurusan >= 9 || len(clanggota.Angkatan) >= 5 || len(clanggota.NoTlp) >= 15 || clanggota.Fakultas >= 3 || clanggota.JKelamin >= 3 {
		var errorMessage string

		if len(clanggota.Nama) >= 50 {
			errorMessage = "Nama tidak boleh lebih dari 50 karakter"
		} else if len(clanggota.Email) >= 50 {
			errorMessage = "Email tidak boleh lebih dari 50 karakter"
		} else if len(clanggota.Nim) >= 11 {
			errorMessage = "Nim tidak boleh lebih dari 11 karakter"
		} else if clanggota.Jurusan >= 9 {
			errorMessage = "Jurusan tidak boleh lebih dari 9"
		} else if len(clanggota.Angkatan) >= 5 {
			errorMessage = "Angkatan tidak boleh lebih dari 5 karakter"
		} else if len(clanggota.NoTlp) >= 15 {
			errorMessage = "Nomor Telepon tidak boleh lebih dari 15 karakter"
		} else if clanggota.Fakultas >= 3 {
			errorMessage = "Fakultas tidak boleh lebih dari 3"
		} else if clanggota.JKelamin >= 3 {
			errorMessage = "Jenis Kelamin tidak boleh lebih dari 3"
		} else {
			errorMessage = ""
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    errorMessage,
		})
		return
	}

	if (len(clanggota.Nama) == 0 || clanggota.Nama == "" || clanggota.Nama == " ") || (len(clanggota.Email) == 0 || clanggota.Email == "" || clanggota.Email == " ") || (len(clanggota.Nim) == 0 || clanggota.Nama == "" || clanggota.Nim == " ") || clanggota.Jurusan == 0 || (len(clanggota.Angkatan) == 0 || clanggota.Angkatan == "" || clanggota.Angkatan == " ") || (len(clanggota.NoTlp) == 0 || clanggota.NoTlp == "" || clanggota.NoTlp == " ") || clanggota.Fakultas == 0 || clanggota.JKelamin == 0 {
		var errorMessage string

		if len(clanggota.Nama) == 0 || clanggota.Nama == "" || clanggota.Nama == " " {
			errorMessage = "Nama tidak boleh kosong"
		} else if len(clanggota.Email) == 0 || clanggota.Email == "" || clanggota.Email == " " {
			errorMessage = "Email tidak boleh kosong"
		} else if len(clanggota.Nim) == 0 || clanggota.Nim == "" || clanggota.Nim == " " {
			errorMessage = "Nim tidak boleh kosong"
		} else if clanggota.Jurusan == 0 {
			errorMessage = "Jurusan tidak boleh kosong"
		} else if len(clanggota.Angkatan) == 0 || clanggota.Angkatan == "" || clanggota.Angkatan == " " {
			errorMessage = "Angkatan tidak boleh kosong"
		} else if len(clanggota.NoTlp) == 0 || clanggota.NoTlp == "" || clanggota.NoTlp == " " {
			errorMessage = "Nomor Telepon tidak kosong"
		} else if clanggota.Fakultas == 0 {
			errorMessage = "Fakultas tidak boleh kosong"
		} else if clanggota.JKelamin == 0 {
			errorMessage = "Jenis Kelamin tidak boleh kosong"
		} else {
			errorMessage = ""
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    errorMessage,
		})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Image tidak boleh kosong",
		})
		return
	}

	// key := c.MustGet("currentUserId").(int)
	key := 1

	// file image
	number := ca.CAUsecase.GenerateID()
	filename := fmt.Sprintf("%s-%s_%s", number, clanggota.Nim, file.Filename)
	clanggota.Img = filename

	err = ca.CAUsecase.RegisterCA(clanggota, key)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	destination := "uploads/image/ca/2023/"
	if err := os.MkdirAll(destination, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create upload directory"})
		return
	}

	// Generate a unique filename or use the original filename
	if err := c.SaveUploadedFile(file, destination+filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	res := model.Response{StatusCode: http.StatusCreated, Message: "Create Calon Anggota Success"}
	c.JSON(http.StatusCreated, res)
}

func (ca *CAController) UpadateCA(c *gin.Context) {
	var clanggota model.RegCA

	if err := c.ShouldBind(&clanggota); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "field is invalid",
		})
		return
	}

	if len(clanggota.Nama) >= 50 || len(clanggota.Email) >= 50 || len(clanggota.Nim) >= 11 || clanggota.Jurusan >= 9 || len(clanggota.Angkatan) >= 5 || len(clanggota.NoTlp) >= 15 || clanggota.Fakultas >= 3 || clanggota.JKelamin >= 3 {
		var errorMessage string

		if len(clanggota.Nama) >= 50 {
			errorMessage = "Nama tidak boleh lebih dari 50 karakter"
		} else if len(clanggota.Email) >= 50 {
			errorMessage = "Email tidak boleh lebih dari 50 karakter"
		} else if len(clanggota.Nim) >= 11 {
			errorMessage = "Nim tidak boleh lebih dari 11 karakter"
		} else if clanggota.Jurusan >= 9 {
			errorMessage = "Jurusan tidak boleh lebih dari 9"
		} else if len(clanggota.Angkatan) >= 5 {
			errorMessage = "Angkatan tidak boleh lebih dari 5 karakter"
		} else if len(clanggota.NoTlp) >= 15 {
			errorMessage = "Nomor Telepon tidak boleh lebih dari 15 karakter"
		} else if clanggota.Fakultas >= 3 {
			errorMessage = "Fakultas tidak boleh lebih dari 3"
		} else if clanggota.JKelamin >= 3 {
			errorMessage = "Jenis Kelamin tidak boleh lebih dari 3"
		} else {
			errorMessage = ""
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    errorMessage,
		})
		return
	}

	if (len(clanggota.Nama) == 0 || clanggota.Nama == "" || clanggota.Nama == " ") || (len(clanggota.Email) == 0 || clanggota.Email == "" || clanggota.Email == " ") || (len(clanggota.Nim) == 0 || clanggota.Nama == "" || clanggota.Nim == " ") || clanggota.Jurusan == 0 || (len(clanggota.Angkatan) == 0 || clanggota.Angkatan == "" || clanggota.Angkatan == " ") || (len(clanggota.NoTlp) == 0 || clanggota.NoTlp == "" || clanggota.NoTlp == " ") || clanggota.Fakultas == 0 || clanggota.JKelamin == 0 {
		var errorMessage string

		if len(clanggota.Nama) == 0 || clanggota.Nama == "" || clanggota.Nama == " " {
			errorMessage = "Nama tidak boleh kosong"
		} else if len(clanggota.Email) == 0 || clanggota.Email == "" || clanggota.Email == " " {
			errorMessage = "Email tidak boleh kosong"
		} else if len(clanggota.Nim) == 0 || clanggota.Nim == "" || clanggota.Nim == " " {
			errorMessage = "Nim tidak boleh kosong"
		} else if clanggota.Jurusan == 0 {
			errorMessage = "Jurusan tidak boleh kosong"
		} else if len(clanggota.Angkatan) == 0 || clanggota.Angkatan == "" || clanggota.Angkatan == " " {
			errorMessage = "Angkatan tidak boleh kosong"
		} else if len(clanggota.NoTlp) == 0 || clanggota.NoTlp == "" || clanggota.NoTlp == " " {
			errorMessage = "Nomor Telepon tidak kosong"
		} else if clanggota.Fakultas == 0 {
			errorMessage = "Fakultas tidak boleh kosong"
		} else if clanggota.JKelamin == 0 {
			errorMessage = "Jenis Kelamin tidak boleh kosong"
		} else {
			errorMessage = ""
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    errorMessage,
		})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Image tidak boleh kosong",
		})
		return
	}

	idCAParam := c.Param("id")
	idCa, err := strconv.Atoi(idCAParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	// key := c.MustGet("currentUserId").(int)
	key := 1

	err = ca.CAUsecase.ValidateID(key)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusNotFound,
			Message:    "Page Not Found",
		})
		return
	}

	// image
	number := ca.CAUsecase.GenerateID()
	filename := fmt.Sprintf("%s-%s_%s", number, clanggota.Nim, file.Filename)
	clanggota.Img = filename

	err = ca.CAUsecase.UpdateCA(clanggota, idCa, key)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	destination := "uploads/image/ca/2023/"
	if err := os.MkdirAll(destination, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create upload directory"})
		return
	}

	// Generate a unique filename or use the original filename
	if err := c.SaveUploadedFile(file, destination+filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	res := model.Response{StatusCode: http.StatusOK, Message: "Update Calon Anggota Success"}
	c.JSON(http.StatusOK, res)
}

func (ca *CAController) ListCA(c *gin.Context) {
	offsetParam := c.Param("offset")
	offsetInt, err := strconv.Atoi(offsetParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}
	limit := c.Param("limit")
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	offset := (offsetInt - 1) * limitInt

	dataCA, err := ca.CAUsecase.ListCA(offset, limitInt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	totalCa, err := ca.CAUsecase.TotalCa()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}
	meta := model.Meta{Offset: offsetInt, Limit: limitInt, Total: totalCa}
	response := model.Response{StatusCode: http.StatusOK, Message: "Success get list calon anggota"}
	result := model.ResponseListCA{Response: response, Meta: meta, Data: dataCA}
	c.JSON(http.StatusOK, result)
}

func (ca *CAController) DeleteCA(c *gin.Context) {
	idCAParam := c.Param("id")
	idCa, err := strconv.Atoi(idCAParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	key := c.MustGet("currentUserId").(int)

	err = ca.CAUsecase.ValidateID(key)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusNotFound,
			Message:    "Page Not Found",
		})
		return
	}

	err = ca.CAUsecase.DeleteCA(idCa, key)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	result := model.Response{StatusCode: http.StatusOK, Message: "Delete Calon Anggota Success"}

	c.JSON(http.StatusOK, result)
}

func (ca *CAController) ImageCa(c *gin.Context) {
	nameImg := c.Param("img")

	file, err := os.Open("uploads/image/ca/2023/" + nameImg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuka gambar"})
		return
	}
	defer file.Close()

	// Mengatur header untuk menunjukkan tipe konten gambar
	c.Header("Content-Type", "image/jpeg")

	// Menyalin isi file gambar ke respons HTTP
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusNotFound,
			Message:    "Page not found",
		})
		return
	}
}
