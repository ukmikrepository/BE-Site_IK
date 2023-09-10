package main

import (
	"backend_ukmik/api/controller"
	"backend_ukmik/api/router"
	"backend_ukmik/config"
	"backend_ukmik/database"
	"backend_ukmik/model"
	"backend_ukmik/repository"
	"backend_ukmik/usecase"
	"log"
	"net/http"
	"time"

	"gorm.io/gorm"
)

func main() {
	loadConfig, err := config.LoadConfigEnv(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	// Database
	db := database.ConnectionDB(&loadConfig)

	errMigrate := db.AutoMigrate(model.User{}, model.CA{})
	if err != nil {
		log.Fatal("🚀 Could not DB Migrate", errMigrate)
	}

	startServer(db)
}

func startServer(db *gorm.DB) {
	AuthenticationRepo := repository.NewAuthenticationRepository(db)
	AuthenticationUseCase := usecase.NewAuthenticationUsecase(AuthenticationRepo)
	AuthenticationController := controller.NewAuthenticationController(AuthenticationUseCase)

	UserRepo := repository.NewUserRepository(db)
	UserUseCase := usecase.NewUserUsecase(UserRepo)
	UserController := controller.NewUserController(UserUseCase)

	CARepo := repository.NewCARepository(db)
	CAUseCase := usecase.NewCAUsecase(CARepo)
	CAController := controller.NewCAController(CAUseCase)

	routes := router.NewRouter(UserRepo, UserController, AuthenticationController, CAController)

	server := &http.Server{
		Addr:           ":3400",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
