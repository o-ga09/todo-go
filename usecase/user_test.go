package usecase_test

import (
	"testing"
	"todo-go/domain"
	mock_port "todo-go/mock/usecase"
	"todo-go/usecase"

	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestUserGetById(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_port.NewMockUserPort(mock)
	mockResult.EXPECT().GetById(1).AnyTimes().Return(domain.User{})
	service := usecase.UserService{mockResult}
	got := service.GetById(1)
	want := domain.User{}
	assert.Equal(t, got, want)
}

func TestUserGetAll(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_port.NewMockUserPort(mock)
	mockResult.EXPECT().GetAll().AnyTimes().Return([]domain.User{})
	service := usecase.UserService{mockResult}
	got := service.GetAll()
	want := []domain.User{}
	assert.Equal(t, got, want)
}

func TestUserCreate(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_port.NewMockUserPort(mock)
	mockResult.EXPECT().Create(domain.UserData{}).AnyTimes().Return(nil)
	service := usecase.UserService{mockResult}
	got := service.Create(domain.UserData{})
	assert.Equal(t, got, nil)
}

func TestUserUpdate(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_port.NewMockUserPort(mock)
	mockResult.EXPECT().Update(1,domain.UserData{}).AnyTimes().Return(nil)
	service := usecase.UserService{mockResult}
	got := service.Update(1,domain.UserData{})
	assert.Equal(t, got, nil)
}

func TestUserDelete(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_port.NewMockUserPort(mock)
	mockResult.EXPECT().Delete(1).AnyTimes().Return(nil)
	service := usecase.UserService{mockResult}
	got := service.Delete(1)
	assert.Equal(t, got, nil)
}