package http

import (
	"GolangProject/internal/user"
	"GolangProject/internal/user/service"
	//HTTP framework
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	response := h.service.GetUsers()
	c.JSON(response.Status, &response)
}

func (h *UserHandler) AddNewUser(c *gin.Context) {
	var u user.User
	response := h.service.AddNewUser(c.ShouldBindJSON(&u), &u)
	c.JSON(response.Status, &response)
}
