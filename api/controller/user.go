package controller

import (
	"backend_ukmik/domain"
	"backend_ukmik/model"
	"fmt"
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

	if err := c.ShouldBind(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "field is invalid" + err.Error(),
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

	key := c.MustGet("currentUserId").(int)
	number := o.UserUsecase.GenerateID()
	filename := fmt.Sprintf("%s-%s_%s", number, user.Name, file.Filename)

	user.Img = filename

	err = o.UserUsecase.CreateUser(user, key)
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
