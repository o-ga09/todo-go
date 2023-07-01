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
	return domain.Task{}
}

func(u *TaskService) GetAll() []domain.Task {
	return []domain.Task{}
}

func(u *TaskService) Create(task domain.TaskData) error {
	return nil
}

func(u *TaskService) Update(id int, task domain.TaskData) error {
	return nil
}

func(u *TaskService) Delete(id int) error {
	return nil
}