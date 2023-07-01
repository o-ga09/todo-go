package handler

import (
	"net/http"
	"todo-go/usecase"

	"github.com/gin-gonic/gin"
)

type ResponseTask struct {
	Taskid int
}

type TaskHandler struct {
	taskService usecase.TaskService
}

func ProviderTaskDriver(taskService usecase.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func(t *TaskHandler) GetTask(c *gin.Context) {
	response := ResponseTask{
		Taskid: 1,
	}
	c.JSON(http.StatusOK,response)
}

func(t *TaskHandler) GetTasks(c *gin.Context) {
	response := ResponseTask{
		Taskid: 1,
	}
	c.JSON(http.StatusOK,response)
}

func(t *TaskHandler) CreateTask(c *gin.Context) {
	response := ResponseTask{
		Taskid: 1,
	}
	c.JSON(http.StatusOK,response)
}

func(t *TaskHandler) UpdateTask(c *gin.Context) {
	response := ResponseTask{
		Taskid: 1,
	}
	c.JSON(http.StatusOK,response)
}

func(t *TaskHandler) DeleteTask(c *gin.Context) {
	response := ResponseTask{
		Taskid: 1,
	}
	c.JSON(http.StatusOK,response)
}