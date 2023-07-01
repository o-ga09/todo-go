package domain

import "time"

type User struct {
	Id int
	Name string
	Password string
	Created_At time.Time
	Update_At time.Time
}

type Task struct {
	Id int
	Name string
	Created_At time.Time
	Update_At time.Time
}

type UserData struct {
	Name string
	Password string
	Created_At time.Time
	Update_At time.Time
}

type TaskData struct {
	Name string
	Created_At time.Time
	Update_At time.Time
}