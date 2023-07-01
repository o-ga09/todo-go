package controller

import (
	di "todo-go/DI"

	"github.com/gin-gonic/gin"
)

func NewServer() (*gin.Engine, error) {
	r := gin.Default()
	// TODO : ミドルウェアで認証周りで使用しているので一旦コメントアウト
	// cfg, err := config.New()
	// if err != nil {
	// 	panic(err)
	// }

	v1 := r.Group("/api/v1")
	{
		userHandler := di.InitUserHandler()
		v1.GET("/user/:id",userHandler.GetUser)
		v1.GET("/users",userHandler.GetUsers)
		v1.POST("/user",userHandler.CreateUser)
		v1.PATCH("/user/:id",userHandler.UpdateUser)
		v1.DELETE("/user/:id",userHandler.DeleteUser)

		taskHandler := di.InitTaskHandler()
		v1.GET("/task/:id",taskHandler.GetTask)
		v1.GET("/tasks",taskHandler.GetTasks)
		v1.POST("/task",taskHandler.CreateTask)
		v1.PATCH("/task/:id",taskHandler.UpdateTask)
		v1.DELETE("/task/:id",taskHandler.DeleteTask)
	}

	return r, nil
}