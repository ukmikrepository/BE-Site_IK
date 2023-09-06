package controller

import (
	"backend_ukmik/domain"
	"backend_ukmik/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

func NewUserController(UserUsecase domain.UserUsecase) *UserController {
	return &UserController{UserUsecase}
}

func (o *UserController) CreateUser(c *gin.Context) {
	var user model.JSONUser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "field json is invalid",
		})
		return
	}

	err := o.UserUsecase.CreateUser(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}
	res := model.Response{StatusCode: http.StatusCreated, Message: "Create User Success"}
	c.JSON(http.StatusCreated, res)
}
