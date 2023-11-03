package handler

import (
	"net/http"
	"strconv"
	"time"
	"todo-go/domain"
	"todo-go/usecase"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskService usecase.TaskService
}

func ProviderTaskDriver(taskService usecase.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func(t *TaskHandler) GetTask(c *gin.Context) {
	rid := c.Param("id")
	id, _ := strconv.Atoi(rid)

	res := t.taskService.TaskInterface.GetById(id)

	response := ResponseTask{
		Taskid: res.Id,
		Name: res.Name,
		CreatedAt: res.Created_At,
		UpdatedAt: res.Update_At,
	}
	c.JSON(http.StatusOK,response)
}

func(t *TaskHandler) GetTasks(c *gin.Context) {
	res := t.taskService.TaskInterface.GetAll()
	response := []ResponseTask{}

	for _, record := range res {
		r := ResponseTask{
			Taskid: record.Id,
			Name: record.Name,
			CreatedAt: record.Created_At,
			UpdatedAt: record.Update_At,
		}

		response = append(response, r)
	}
	c.JSON(http.StatusOK,response)
}

func(t *TaskHandler) CreateTask(c *gin.Context) {
	rname := c.PostForm("name")

	request := domain.TaskData{
		Name: rname,
		Created_At: time.Now(),
		Update_At: time.Now(),
	}
	err := t.taskService.TaskInterface.Create(request)

	if err != nil {
		c.JSON(http.StatusBadRequest,MessageResponse{Message: err.Error()})	
		return
	}
	c.JSON(http.StatusOK,MessageResponse{Message: "registerd"})
}

func(t *TaskHandler) UpdateTask(c *gin.Context) {
	rid := c.Param("id")
	rname := c.PostForm("name")

	id, _ := strconv.Atoi(rid)

	request := domain.TaskData{
		Name: rname,
		Created_At: time.Now(),
		Update_At: time.Now(),
	}
	err := t.taskService.TaskInterface.Update(id,request)

	if err != nil {
		c.JSON(http.StatusBadRequest,MessageResponse{Message: err.Error()})	
		return
	}
	c.JSON(http.StatusOK,MessageResponse{Message: "updated"})
}

func(t *TaskHandler) DeleteTask(c *gin.Context) {
	rid := c.Param("id")
	id, _ := strconv.Atoi(rid)
	
	err := t.taskService.TaskInterface.Delete(id)

	if err != nil {
		c.JSON(http.StatusOK,MessageResponse{Message: err.Error()})	
	}
	c.JSON(http.StatusOK,"deleted")
}

type ResponseTask struct {
	Taskid    int       `json:"taskid"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RequestTask struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}