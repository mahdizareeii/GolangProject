package main

import (
	"GolangProject/pkg/db"

	handler "GolangProject/internal/user/delivery/http"
	userRepo "GolangProject/internal/user/repository"
	service "GolangProject/internal/user/service"

	//HTTP framework
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	dataBase := db.GetDbConnection()
	userRepository := userRepo.CreateUserRepository(dataBase)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	r.GET("/users", userHandler.GetUsers)
	r.POST("/user", userHandler.AddNewUser)
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
