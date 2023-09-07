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

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image upload is required"})
		return
	}

	key := c.MustGet("currentUserId").(int)

	// file image
	filename := fmt.Sprintf("%s_%s", clanggota.Nim, file.Filename)
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

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image upload is required"})
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

	key := c.MustGet("currentUserId").(int)

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
	filename := fmt.Sprintf("%s_%s", clanggota.Nim, file.Filename)
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengirim gambar"})
	}
}
