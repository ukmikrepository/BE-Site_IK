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

	// err := a.UserUsecase.CreateUser(user)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
	// 		StatusCode: http.StatusInternalServerError,
	// 		Message:    err.Error(),
	// 	})
	// 	return
	// }
	// res := model.Response{StatusCode: http.StatusCreated, Message: "Create User Success"}
	c.JSON(http.StatusCreated, login)
}
