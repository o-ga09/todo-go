package main

import "todo-go/controller"

func main() {
	server, err := controller.NewServer()

	if err != nil {
		panic(err)
	}
	server.Run()
}