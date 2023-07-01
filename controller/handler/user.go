package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseUser struct {
	Userid int
}

type UserHnadler struct {}

func(h *UserHnadler) GetUser(c *gin.Context) {
	response := ResponseUser{
		Userid: 1,
	}
	c.JSON(http.StatusOK,response)
}

func(h *UserHnadler) GetUsers(c *gin.Context) {
	response := ResponseUser{
		Userid: 1,
	}
	c.JSON(http.StatusOK,response)
}

func(h *UserHnadler) CreateUser(c *gin.Context) {
	response := ResponseUser{
		Userid: 1,
	}
	c.JSON(http.StatusOK,response)
}

func(h *UserHnadler) UpdateUser(c *gin.Context) {
	response := ResponseUser{
		Userid: 1,
	}
	c.JSON(http.StatusOK,response)
}

func(h *UserHnadler) DeleteUser(c *gin.Context) {
	response := ResponseUser{
		Userid: 1,
	}
	c.JSON(http.StatusOK,response)
}