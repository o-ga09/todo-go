package di

import (
	"todo-go/controller/handler"
	"todo-go/driver"
	"todo-go/gateway"
	"todo-go/usecase"
)

func InitUserHandler() *handler.UserHandler {
	db, err := driver.New()
	if err != nil {
		panic(err)
	}
	driver := driver.ProviderUserDriver(db)
	gateway := gateway.ProviderUserDriver(driver)
	usecase := usecase.ProviderUserdriver(gateway)
	handler := handler.ProviderUserDriver(usecase)

	return handler
}

func InitTaskHandler() *handler.TaskHandler {
	db, err := driver.New()
	if err != nil {
		panic(err)
	}
	driver := driver.ProviderTaskDriver(db)
	gateway := gateway.ProviderTaskDriver(driver)
	usecase := usecase.ProviderTaskDriver(gateway)
	handler := handler.ProviderTaskDriver(usecase)
	
	return handler
}