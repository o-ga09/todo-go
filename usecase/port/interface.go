package port

import "todo-go/domain"

type UserPort interface {
	GetById(int) domain.User
	GetAll() []domain.User
	Create(domain.UserData) error 
	Update(id int, user domain.UserData) error
	Delete(int) error
}

type TaskPort interface {
	GetById(int) domain.Task
	GetAll() []domain.Task
	Create(domain.TaskData) error 
	Update(id int, user domain.TaskData) error
	Delete(int) error
}