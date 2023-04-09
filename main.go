package main

import (
	"net/http"

	"github.com/SyaibanAhmadRamadhan/impl-repo-pattern-go/config"
	"github.com/SyaibanAhmadRamadhan/impl-repo-pattern-go/controllers"
	"github.com/SyaibanAhmadRamadhan/impl-repo-pattern-go/exception"
	"github.com/SyaibanAhmadRamadhan/impl-repo-pattern-go/repositories"
	"github.com/SyaibanAhmadRamadhan/impl-repo-pattern-go/services"
	
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	db := config.ConnectDb()

	userRepo := repositories.NewUserRepositoryImpl()
	userService := services.NewUserServiceImpl(userRepo, db, validator.New())
	userHandler := controllers.NewUserControllerImpl(userService)

	router := httprouter.New()
	router.POST("/api/user", userHandler.Create)
	router.GET("/api/user", userHandler.FindAll)
	router.GET("/api/user/:id/detail", userHandler.FindById)
	router.PUT("/api/user/:id/update", userHandler.Update)
	router.DELETE("/api/user/:id/delete", userHandler.Delete)

	router.PanicHandler = exception.ErrorHandler
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
