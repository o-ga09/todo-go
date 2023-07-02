package driver

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type UserDriver interface {
	GetById(int) User
	GetAll() []User
	Create(UserData) error 
	Update(id int, user UserData) error
	Delete(int) error
}

type User struct {
	Id        int       `gorm:"primaryKey"`
	Name      string    `gorm:"size:255"`
	Password  string	`gorm:"size:255"`
	CreatedAt time.Time `gorm:"type:datetime"`
	UpdatedAt time.Time `gorm:"type:datetime"`
}

type UserData struct {
	Name      string    `json:"name"`
	Password  string	`json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserDriverImpl struct {
	conn *gorm.DB
}

func ProviderUserDriver(conn *gorm.DB) UserDriver {
	return &UserDriverImpl{conn: conn}
}

func(d *UserDriverImpl) GetById(id int) User {
	user := User{}
	d.conn.First(&user,id)
	return user
}

func(d *UserDriverImpl) GetAll() []User {
	users := []User{}
	d.conn.Find(&users)
	return users
}

func(d *UserDriverImpl) Create(user UserData) error {
	err := d.conn.Create(&user)
	log.Print(err)
	return err.Error
}

func(d *UserDriverImpl) Update(id int, user UserData) error {
	err := d.conn.Model(&User{}).Where("id = ?", id).Updates(&user)
	return err.Error
}

func(d *UserDriverImpl) Delete(id int) error {
	err := d.conn.Delete(&User{},id)
	return err.Error
}

func (UserData) TableName() string {
	return "user"
}