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
	res := u.UserInterface.GetById(id)
	return res
}

func(u *UserService) GetAll() []domain.User {
	res := u.UserInterface.GetAll()
	return res
}

func(u *UserService) Create(user domain.UserData) error {
	err := u.UserInterface.Create(user)
	return err
}

func(u *UserService) Update(id int, user domain.UserData) error {
	err := u.UserInterface.Update(id,user)
	return err
}

func(u *UserService) Delete(id int) error {
	err := u.UserInterface.Delete(id)
	return err
}