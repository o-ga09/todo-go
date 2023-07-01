package gateway_test

import (
	"testing"
	"todo-go/domain"
	"todo-go/driver"
	"todo-go/gateway"
	mock_driver "todo-go/mock/userDriver"

	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestUserGetById(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_driver.NewMockUserDriver(mock)
	mockResult.EXPECT().GetById(1).AnyTimes().Return(driver.User{})
	gatewayDriver := gateway.ProviderUserDriver(mockResult)
	got := gatewayDriver.GetById(1)
	want := domain.User{}
	assert.Equal(t, got, want)
}

func TestUserGetALL(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_driver.NewMockUserDriver(mock)
	mockResult.EXPECT().GetAll().AnyTimes().Return([]driver.User{})
	gatewayDriver := gateway.ProviderUserDriver(mockResult)
	got := gatewayDriver.GetAll()
	want := []domain.User{}
	assert.Equal(t, got, want)
}

func TestUserCreate(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_driver.NewMockUserDriver(mock)
	mockResult.EXPECT().Create(driver.UserData{}).AnyTimes().Return(nil)
	gatewayDriver := gateway.ProviderUserDriver(mockResult)
	got := gatewayDriver.Create(domain.UserData{})
	assert.Equal(t, got, nil)
}

func TestUserUpdate(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_driver.NewMockUserDriver(mock)
	mockResult.EXPECT().Update(1,driver.UserData{}).AnyTimes().Return(nil)
	gatewayDriver := gateway.ProviderUserDriver(mockResult)
	got := gatewayDriver.Update(1,domain.UserData{})
	assert.Equal(t, got, nil)
}

func TestUserDelete(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_driver.NewMockUserDriver(mock)
	mockResult.EXPECT().Delete(1).AnyTimes().Return(nil)
	gatewayDriver := gateway.ProviderUserDriver(mockResult)
	got := gatewayDriver.Delete(1)
	assert.Equal(t, got, nil)
}