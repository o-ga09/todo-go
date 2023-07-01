package gateway

import (
	"todo-go/domain"
	"todo-go/driver"
)

type UserGateway struct {
	userDriver driver.UserDriver
}

func ProviderUserDriver(userDriver driver.UserDriver) *UserGateway {
	return &UserGateway{userDriver: userDriver}
}

func(g *UserGateway) GetById(id int) domain.User {
	res := g.userDriver.GetById(id)

	result := domain.User{
		Id: res.Id,
		Name: res.Name,
		Password: res.Password,
		Created_At: res.CreatedAt,
		Update_At: res.UpdatedAt,
	}
	return result
}

func(g *UserGateway) GetAll() []domain.User {
	res := g.userDriver.GetAll()
	result := []domain.User{}
	for _, record := range res {
		user := domain.User{
			Id: record.Id,
			Name: record.Name,
			Password: record.Password,
			Created_At: record.CreatedAt,
			Update_At: record.UpdatedAt,
		}

		result = append(result,user)
	}
	return result
}

func(g *UserGateway) Create(user domain.UserData) error {
	driverUser := driver.UserData{
		Name: user.Name,
		Password: user.Password,
		CreatedAt: user.Created_At,
		UpdatedAt: user.Update_At,
	}
	err := g.userDriver.Create(driverUser)
	return err
}

func(g *UserGateway) Update(id int, user domain.UserData) error {
	driverUser := driver.UserData{
		Name: user.Name,
		Password: user.Password,
		CreatedAt: user.Created_At,
		UpdatedAt: user.Update_At,
	}
	err := g.userDriver.Update(id,driverUser)
	return err
}

func(g *UserGateway) Delete(id int) error {
	err := g.userDriver.Delete(id)
	return err
}