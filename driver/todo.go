package driver

import (
	"time"

	"gorm.io/gorm"
)

type TaskDriver interface {
	GetById(int) Task
	GetAll() []Task
	Create(TaskData) error 
	Update(id int, task TaskData) error
	Delete(int) error
}

type Task struct {
	Id        int       `gorm:"primaryKey"`
	Name      string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"type:datetime"`
	UpdatedAt time.Time `gorm:"type:datetime"`
}

type TaskData struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TaskDriverImpl struct {
	conn *gorm.DB
}

func ProviderTaskDriver(conn *gorm.DB) TaskDriver {
	return &TaskDriverImpl{conn: conn}
}

func(t *TaskDriverImpl) GetById(id int) Task {
	task := Task{}
	t.conn.First(&task,id)
	return task
}

func(t *TaskDriverImpl) GetAll() []Task {
	tasks := []Task{}
	t.conn.Find(&tasks)
	return tasks
}

func(t *TaskDriverImpl) Create(task TaskData) error {
	err := t.conn.Create(&task)
	return err.Error
}

func(t *TaskDriverImpl) Update(id int, task TaskData) error {
	err := t.conn.Model(&Task{}).Where("id = ?", id).Updates(&task)
	return err.Error
}

func(t *TaskDriverImpl) Delete(id int) error {
	err := t.conn.Delete(&Task{},id)
	return err.Error
}

func(Task) TableName() string {
	return "task"
}

func(TaskData) TableName() string {
	return "task"
}