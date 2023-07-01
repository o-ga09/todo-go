package handler

import (
	"net/http"
	"todo-go/usecase"

	"github.com/gin-gonic/gin"
)

type ResponseUser struct {
	Userid int
}

type UserHandler struct {
	userService usecase.UserService
}

func ProviderUserDriver(userService *usecase.UserService) *UserHandler {
	return &UserHandler{userService: *userService}
}

func(h *UserHandler) GetUser(c *gin.Context) {
	response := ResponseUser{
		Userid: 1,
	}
	c.JSON(http.StatusOK,response)
}

func(h *UserHandler) GetUsers(c *gin.Context) {
	response := ResponseUser{
		Userid: 1,
	}
	c.JSON(http.StatusOK,response)
}

func(h *UserHandler) CreateUser(c *gin.Context) {
	response := ResponseUser{
		Userid: 1,
	}
	c.JSON(http.StatusOK,response)
}

func(h *UserHandler) UpdateUser(c *gin.Context) {
	response := ResponseUser{
		Userid: 1,
	}
	c.JSON(http.StatusOK,response)
}

func(h *UserHandler) DeleteUser(c *gin.Context) {
	response := ResponseUser{
		Userid: 1,
	}
	c.JSON(http.StatusOK,response)
}