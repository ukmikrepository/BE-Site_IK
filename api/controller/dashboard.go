package controller

import (
	"backend_ukmik/domain"
	"backend_ukmik/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DashboardController struct {
	DashboardUsecase domain.DashboardUsecase
}

func NewDashboardController(DashboardUsecase domain.DashboardUsecase) *DashboardController {
	return &DashboardController{DashboardUsecase}
}

func (d *DashboardController) Dashboard(c *gin.Context) {
	key := c.MustGet("currentUserId").(int)
	_, err := d.DashboardUsecase.Dashboarad(key)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	c.JSON(200, "dashboard")
}
