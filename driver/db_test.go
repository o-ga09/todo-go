package driver_test

import (
	"testing"
	"todo-go/driver"

	"github.com/magiconair/properties/assert"
)

func TestConnectDB(t *testing.T) {
	t.Setenv("ENV","dev")
	t.Setenv("DBUSER","todo")
	t.Setenv("DBPASSWORD","password")
	t.Setenv("DBHOST","127.0.0.1")
	t.Setenv("DBNAME","todo")
	t.Setenv("DBPORT","3306")

	_ , err := driver.New()
	assert.Equal(t,nil,err)
}