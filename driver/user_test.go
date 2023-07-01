package driver_test

import (
	"testing"
	"todo-go/driver"
	mock_driver "todo-go/mock/userDriver"

	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestGetUserById(t *testing.T) {
	t.Setenv("ENV","dev")
	t.Setenv("DBUSER","todo")
	t.Setenv("DBPASSWORD","password")
	t.Setenv("DBHOST","127.0.0.1")
	t.Setenv("DBNAME","todo")
	t.Setenv("DBPORT","3306")

	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_driver.NewMockUserDriver(mock)
	mockResult.EXPECT().GetById(1).AnyTimes().Return(driver.User{})
	conn ,_ := driver.New()

	driverImpl := driver.ProviderUserDriver(conn)
	got := driverImpl.GetById(1)
	assert.Equal(t, got, driver.User{})
}

func TestGetUserAll(t *testing.T) {
	t.Setenv("ENV","dev")
	t.Setenv("DBUSER","todo")
	t.Setenv("DBPASSWORD","password")
	t.Setenv("DBHOST","127.0.0.1")
	t.Setenv("DBNAME","todo")
	t.Setenv("DBPORT","3306")

	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_driver.NewMockUserDriver(mock)
	mockResult.EXPECT().GetAll().AnyTimes().Return([]driver.User{})
	conn ,_ := driver.New()

	driverImpl := driver.ProviderUserDriver(conn)
	got := driverImpl.GetAll()
	assert.Equal(t, got, []driver.User{})
}

func TestCreateUser(t *testing.T) {
	t.Setenv("ENV","dev")
	t.Setenv("DBUSER","todo")
	t.Setenv("DBPASSWORD","password")
	t.Setenv("DBHOST","127.0.0.1")
	t.Setenv("DBNAME","todo")
	t.Setenv("DBPORT","3306")

	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_driver.NewMockUserDriver(mock)
	mockResult.EXPECT().Create(driver.User{}).AnyTimes().Return(nil)
	conn ,_ := driver.New()

	driverImpl := driver.ProviderUserDriver(conn)
	got := driverImpl.Create(driver.UserData{})
	assert.Equal(t, got, nil)
}

func TestUpdateUser(t *testing.T) {
	t.Setenv("ENV","dev")
	t.Setenv("DBUSER","todo")
	t.Setenv("DBPASSWORD","password")
	t.Setenv("DBHOST","127.0.0.1")
	t.Setenv("DBNAME","todo")
	t.Setenv("DBPORT","3306")

	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_driver.NewMockUserDriver(mock)
	mockResult.EXPECT().Update(1,driver.UserData{}).AnyTimes().Return(nil)
	conn ,_ := driver.New()

	driverImpl := driver.ProviderUserDriver(conn)
	got := driverImpl.Update(1,driver.UserData{})
	assert.Equal(t, got, nil)
}

func TestDeleteUser(t *testing.T) {
	t.Setenv("ENV","dev")
	t.Setenv("DBUSER","todo")
	t.Setenv("DBPASSWORD","password")
	t.Setenv("DBHOST","127.0.0.1")
	t.Setenv("DBNAME","todo")
	t.Setenv("DBPORT","3306")

	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_driver.NewMockUserDriver(mock)
	mockResult.EXPECT().Delete(1).AnyTimes().Return(nil)
	conn ,_ := driver.New()

	driverImpl := driver.ProviderUserDriver(conn)
	got := driverImpl.Delete(1)
	assert.Equal(t, got, nil)
}