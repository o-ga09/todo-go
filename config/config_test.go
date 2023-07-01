package config_test

import (
	"testing"
	"todo-go/config"

	"github.com/magiconair/properties/assert"
)

func TestConfig(t *testing.T) {
	t.Setenv("ENV","dev")
	t.Setenv("DBUSER","todo")
	t.Setenv("DBPASSWORD","password")
	t.Setenv("DBHOST","dbprod01")
	t.Setenv("DBNAME","todo")
	t.Setenv("DBPORT","3306")

	expected := &config.Config{
		Env: "dev",
		DBUser: "todo",
		DBPassword: "password",
		DBHost: "dbprod01",
		DBName: "todo",
		DBPort: "3306",
	}
	got, err := config.New()
	assert.Equal(t,expected, got)
	assert.Equal(t, nil, err)
}