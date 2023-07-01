package config

import (
	"github.com/caarlos0/env"
)

type Config struct {
	Env        string `env:"ENV" json:"env"`
	DBHost     string `env:"DBHOST" json:"db_host"`
	DBUser     string `env:"DBUSER" json:"db_user"`
	DBPassword string `env:"DBPASSWORD" json:"db_password"`
	DBName     string `env:"DBNAME" json:"db_name"`
	DBPort     string `env:"DBPORT" json:"db_port"`
}
func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg, nil
}