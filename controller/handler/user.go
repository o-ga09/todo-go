package handler

import (
	"net/http"
	"strconv"
	"time"
	"todo-go/domain"
	"todo-go/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService usecase.UserService
}

func ProviderUserDriver(userService *usecase.UserService) *UserHandler {
	return &UserHandler{userService: *userService}
}

func(h *UserHandler) GetUser(c *gin.Context) {
	rid := c.Param("id")
	id, _ := strconv.Atoi(rid)

	res := h.userService.UserInterface.GetById(id)

	response := ResponseUser{
		Userid: res.Id,
		Name: res.Name,
		Password: res.Password,
		CreatedAt: res.Created_At,
		UpdatedAt: res.Update_At,
	}
	c.JSON(http.StatusOK,response)
}

func(h *UserHandler) GetUsers(c *gin.Context) {
	res := h.userService.UserInterface.GetAll()
	response := []ResponseUser{}

	for _, record := range res {
		r := ResponseUser{
			Userid: record.Id,
			Name: record.Name,
			Password: record.Password,
			CreatedAt: record.Created_At,
			UpdatedAt: record.Update_At,
		}

		response = append(response, r)
	}
	c.JSON(http.StatusOK,response)
}

func(h *UserHandler) CreateUser(c *gin.Context) {
	rname := c.PostForm("name")
	rpassword := c.PostForm("password")

	request := domain.UserData{
		Name: rname,
		Password: rpassword,
		Created_At: time.Now(),
		Update_At: time.Now(),
	}
	err := h.userService.UserInterface.Create(request)

	if err != nil {
		c.JSON(http.StatusBadRequest,MessageResponse{Message: err.Error()})	
		return
	}
	c.JSON(http.StatusOK,MessageResponse{Message: "registerd"})
}

func(h *UserHandler) UpdateUser(c *gin.Context) {
	rid := c.Param("id")
	rname := c.PostForm("name")
	rpassword := c.PostForm("password")

	id, _ := strconv.Atoi(rid)

	request := domain.UserData{
		Name: rname,
		Password: rpassword,
		Created_At: time.Now(),
		Update_At: time.Now(),
	}
	err := h.userService.UserInterface.Update(id,request)

	if err != nil {
		c.JSON(http.StatusBadRequest,MessageResponse{Message: err.Error()})	
		return
	}
	c.JSON(http.StatusOK,MessageResponse{Message: "updated"})
}

func(h *UserHandler) DeleteUser(c *gin.Context) {
	rid := c.Param("id")
	id, _ := strconv.Atoi(rid)
	
	err := h.userService.UserInterface.Delete(id)

	if err != nil {
		c.JSON(http.StatusOK,MessageResponse{Message: err.Error()})	
	}
	c.JSON(http.StatusOK,"deleted")
}

type ResponseUser struct {
	Userid    int       `json:"userid"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RequestUser struct {
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MessageResponse struct {
	Message string `json:"message"`
}
