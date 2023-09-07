package router

import (
	"backend_ukmik/api/controller"
	"backend_ukmik/api/middleware"
	"backend_ukmik/domain"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(userDomain domain.UserRepository, UserController *controller.UserController, AuthenticationController *controller.AuthenticationController, CAController *controller.CAController) *gin.Engine {
	service := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://10.0.26.57:3000", "*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Authorization", "Content-Type", "Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With"}
	config.AllowCredentials = true

	service.Use(cors.New(config))

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")

	// authentication
	router.POST("/login", AuthenticationController.Login)

	// test
	router.POST("/crate_user", UserController.CreateUser)

	// pendaftaran calon anggota
	router.POST("/ca-register", middleware.DeserializeAdminRole(userDomain), CAController.RegisterCA)
	router.PUT("/ca-update/:id", middleware.DeserializeAdminRole(userDomain), CAController.UpadateCA)
	router.GET("/ca-list/:offset/:limit", middleware.DeserializeAdminRole(userDomain), CAController.ListCA)
	router.DELETE("/ca-delete/:id", middleware.DeserializeAdminRole(userDomain), CAController.DeleteCA)
	router.GET("/ca-image/:img", middleware.DeserializeAdminRole(userDomain), CAController.ImageCa)

	return service
}
