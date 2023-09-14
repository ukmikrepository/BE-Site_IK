package controller

import (
	"backend_ukmik/domain"
	"backend_ukmik/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	AuthenticationUsecase domain.AuthenticationUsecase
}

func NewAuthenticationController(AuthenticationUsecase domain.AuthenticationUsecase) *AuthenticationController {
	return &AuthenticationController{AuthenticationUsecase}
}

func (a *AuthenticationController) Login(c *gin.Context) {
	var login model.Login

	if err := c.ShouldBindJSON(&login); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "field json is invalid",
		})
		return
	}

	token, err := a.AuthenticationUsecase.ValidateUser(login)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	res := model.Response{StatusCode: http.StatusOK, Message: "Login successful"}
	result := model.ResLogin{Res: res, Token: token}
	c.JSON(res.StatusCode, result)
}

func (a *AuthenticationController) Logout(c *gin.Context) {
	// Di sini Anda dapat mengimplementasikan logika logout sesuai dengan kebutuhan Anda,
	// seperti menghapus token sesi atau tindakan logout lainnya.

	// Contoh sederhana: hanya mengirimkan pesan logout berhasil.
	c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Logout berhasil",
	})
}
