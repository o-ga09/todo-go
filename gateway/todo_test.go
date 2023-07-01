package gateway_test

import (
	"testing"
	"todo-go/domain"
	"todo-go/driver"
	"todo-go/gateway"
	mock_driver "todo-go/mock/todoDriver"

	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestTaskGetById(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_driver.NewMockTaskDriver(mock)
	gatewayDriver := gateway.ProviderTaskDriver(mockResult)
	mockResult.EXPECT().GetById(1).AnyTimes().Return(driver.Task{})
	got := gatewayDriver.GetById(1)
	want := domain.Task{}
	assert.Equal(t,got,want)
}

func TestTaskGetAll(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_driver.NewMockTaskDriver(mock)
	gatewayDriver := gateway.ProviderTaskDriver(mockResult)
	mockResult.EXPECT().GetAll().AnyTimes().Return([]driver.Task{})
	got := gatewayDriver.GetAll()
	want := []domain.Task{}
	assert.Equal(t,got,want)
}

func TestTaskCreate(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_driver.NewMockTaskDriver(mock)
	gatewayDriver := gateway.ProviderTaskDriver(mockResult)
	mockResult.EXPECT().Create(driver.TaskData{}).AnyTimes().Return(nil)
	got := gatewayDriver.Create(domain.TaskData{})
	assert.Equal(t,got,nil)
}

func TestTaskUpdate(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_driver.NewMockTaskDriver(mock)
	gatewayDriver := gateway.ProviderTaskDriver(mockResult)
	mockResult.EXPECT().Update(1,driver.TaskData{}).AnyTimes().Return(nil)
	got := gatewayDriver.Update(1,domain.TaskData{})
	assert.Equal(t,got,nil)
}

func TestTaskDelete(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_driver.NewMockTaskDriver(mock)
	gatewayDriver := gateway.ProviderTaskDriver(mockResult)
	mockResult.EXPECT().Delete(1).AnyTimes().Return(nil)
	got := gatewayDriver.Delete(1)
	assert.Equal(t,got,nil)
}