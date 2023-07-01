package usecase

import (
	"todo-go/domain"
	"todo-go/usecase/port"
)

type UserService struct {
	UserInterface port.UserPort
}

func ProviderUserdriver(UserGateway port.UserPort) *UserService {
	return &UserService{UserInterface: UserGateway}
}

func(u *UserService) GetById(id int) domain.User {
	return domain.User{}
}

func(u *UserService) GetAll() []domain.User {
	return []domain.User{}
}

func(u *UserService) Create(user domain.UserData) error {
	return nil
}

func(u *UserService) Update(id int, user domain.UserData) error {
	return nil
}

func(u *UserService) Delete(id int) error {
	return nil
}