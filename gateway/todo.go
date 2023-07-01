package gateway

import (
	"todo-go/domain"
	"todo-go/driver"
)

type TaskGateway struct {
	taskGateway driver.TaskDriver
}

func ProviderTaskDriver(taskDriver driver.TaskDriver) *TaskGateway {
	return &TaskGateway{taskGateway: taskDriver}
}

func(g *TaskGateway) GetById(id int) domain.Task {
	res := g.taskGateway.GetById(id)
	task := domain.Task{
		Id: res.Id,
		Name: res.Name,
		Created_At: res.CreatedAt,
		Update_At: res.UpdatedAt,
	}
	return task
}

func(g *TaskGateway) GetAll() []domain.Task {
	res := g.taskGateway.GetAll()

	result := []domain.Task{}
	for _, record := range res {
		task := domain.Task{
			Id: record.Id,
			Name: record.Name,
			Created_At: record.CreatedAt,
			Update_At: record.UpdatedAt,
		}
		result = append(result,task)
	}
	return result
}

func(g *TaskGateway) Create(task domain.TaskData) error {
	driverTask := driver.TaskData{
		Name: task.Name,
		CreatedAt: task.Created_At,
		UpdatedAt: task.Update_At,
	}
	err := g.taskGateway.Create(driverTask)
	return err
}

func(g *TaskGateway) Update(id int, task domain.TaskData) error {
	driverTask := driver.TaskData{
		Name: task.Name,
		CreatedAt: task.Created_At,
		UpdatedAt: task.Update_At,
	}
	err := g.taskGateway.Update(id,driverTask)
	return err
}

func(g *TaskGateway) Delete(id int) error {
	err := g.taskGateway.Delete(id)
	return err
}