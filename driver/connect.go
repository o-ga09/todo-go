package driver

import (
	"fmt"
	"time"
	"todo-go/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func New() (*gorm.DB, error) {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	var dialector gorm.Dialector
	if cfg.Env == "PROD" {
		dialector = mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local&tls=true",
		cfg.DBUser, cfg.DBPassword,
		cfg.DBHost, cfg.DBName,
		))
	} else {
		dialector = mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost,
		cfg.DBPort, cfg.DBName,
		))
	}
	
	if db,err = gorm.Open(dialector,&gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}}); err != nil {
		retry(dialector,100)
	}

	return db, nil
}

func retry(dialector gorm.Dialector, count uint) {
	var err error
	if db, err = gorm.Open(dialector,&gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}}); err != nil {
		if count > 1 {
			time.Sleep(time.Second * 2)
			count--
			fmt.Printf("retry... count:%v\n", count)
			retry(dialector, count)
			return
		}
		panic(err.Error())
	}
}