package driver_test

import (
	"testing"
	"todo-go/driver"
	mock_driver "todo-go/mock/todoDriver"

	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestGetTaskById(t *testing.T) {
	t.Setenv("ENV","dev")
	t.Setenv("DBUSER","todo")
	t.Setenv("DBPASSWORD","password")
	t.Setenv("DBHOST","127.0.0.1")
	t.Setenv("DBNAME","todo")
	t.Setenv("DBPORT","3306")

	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_driver.NewMockTaskDriver(mock)
	mockResult.EXPECT().GetById(1).AnyTimes().Return(driver.Task{})
	conn ,_ := driver.New()

	driverImpl := driver.ProviderTaskDriver(conn)
	got := driverImpl.GetById(1)
	assert.Equal(t, got, driver.Task{})
}

func TestGetTaskAll(t *testing.T) {
	t.Setenv("ENV","dev")
	t.Setenv("DBUSER","todo")
	t.Setenv("DBPASSWORD","password")
	t.Setenv("DBHOST","127.0.0.1")
	t.Setenv("DBNAME","todo")
	t.Setenv("DBPORT","3306")

	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_driver.NewMockTaskDriver(mock)
	mockResult.EXPECT().GetAll().AnyTimes().Return([]driver.Task{})
	conn ,_ := driver.New()

	driverImpl := driver.ProviderTaskDriver(conn)
	got := driverImpl.GetAll()
	assert.Equal(t, got, []driver.Task{})
}

func TestCreateTask(t *testing.T) {
	t.Setenv("ENV","dev")
	t.Setenv("DBUSER","todo")
	t.Setenv("DBPASSWORD","password")
	t.Setenv("DBHOST","127.0.0.1")
	t.Setenv("DBNAME","todo")
	t.Setenv("DBPORT","3306")

	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_driver.NewMockTaskDriver(mock)
	mockResult.EXPECT().Create(driver.TaskData{}).AnyTimes().Return(nil)
	conn ,_ := driver.New()

	driverImpl := driver.ProviderTaskDriver(conn)
	got := driverImpl.Create(driver.TaskData{})
	assert.Equal(t, got, nil)
}

func TestUpdateTask(t *testing.T) {
	t.Setenv("ENV","dev")
	t.Setenv("DBUSER","todo")
	t.Setenv("DBPASSWORD","password")
	t.Setenv("DBHOST","127.0.0.1")
	t.Setenv("DBNAME","todo")
	t.Setenv("DBPORT","3306")

	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_driver.NewMockTaskDriver(mock)
	mockResult.EXPECT().Update(1,driver.TaskData{}).AnyTimes().Return(nil)
	conn ,_ := driver.New()

	driverImpl := driver.ProviderTaskDriver(conn)
	got := driverImpl.Update(1,driver.TaskData{})
	assert.Equal(t, got, nil)
}

func TestDeleteTask(t *testing.T) {
	t.Setenv("ENV","dev")
	t.Setenv("DBUSER","todo")
	t.Setenv("DBPASSWORD","password")
	t.Setenv("DBHOST","127.0.0.1")
	t.Setenv("DBNAME","todo")
	t.Setenv("DBPORT","3306")

	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_driver.NewMockTaskDriver(mock)
	mockResult.EXPECT().Delete(1).AnyTimes().Return(nil)
	conn ,_ := driver.New()

	driverImpl := driver.ProviderTaskDriver(conn)
	got := driverImpl.Delete(1)
	assert.Equal(t, got, nil)
}