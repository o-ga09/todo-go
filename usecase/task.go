package usecase

import (
	"todo-go/domain"
	"todo-go/usecase/port"
)

type TaskService struct {
	TaskInterface port.TaskPort
}

func ProviderTaskDriver(TaskGateway port.TaskPort) TaskService {
	return TaskService{TaskInterface: TaskGateway}
}

func(u *TaskService) GetById(id int) domain.Task {
	res := u.TaskInterface.GetById(id)
	return res
}

func(u *TaskService) GetAll() []domain.Task {
	res := u.TaskInterface.GetAll()
	return res
}

func(u *TaskService) Create(task domain.TaskData) error {
	err := u.TaskInterface.Create(task)
	return err
}

func(u *TaskService) Update(id int, task domain.TaskData) error {
	err := u.TaskInterface.Update(id,task)
	return err
}

func(u *TaskService) Delete(id int) error {
	err := u.TaskInterface.Delete(id)
	return err
}