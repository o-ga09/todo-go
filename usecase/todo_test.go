package usecase_test

import (
	"testing"
	"todo-go/domain"
	mock_port "todo-go/mock/usecase"
	"todo-go/usecase"

	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestTaskGetById(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_port.NewMockTaskPort(mock)
	mockResult.EXPECT().GetById(1).AnyTimes().Return(domain.Task{})
	service := usecase.ProviderTaskDriver(mockResult)
	got := service.GetById(1)
	want := domain.Task{}
	assert.Equal(t, got , want )
}

func TestTaskGetAll(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_port.NewMockTaskPort(mock)
	mockResult.EXPECT().GetAll().AnyTimes().Return([]domain.Task{})
	service := usecase.ProviderTaskDriver(mockResult)
	got := service.GetAll()
	want := []domain.Task{}
	assert.Equal(t, got , want )
}

func TestTaskCreate(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_port.NewMockTaskPort(mock)
	mockResult.EXPECT().Create(domain.TaskData{}).AnyTimes().Return(nil)
	service := usecase.ProviderTaskDriver(mockResult)
	got := service.Create(domain.TaskData{})
	assert.Equal(t, got , nil )
}

func TestTaskUpdate(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_port.NewMockTaskPort(mock)
	mockResult.EXPECT().Update(1,domain.TaskData{}).AnyTimes().Return(nil)
	service := usecase.ProviderTaskDriver(mockResult)
	got := service.Update(1,domain.TaskData{})
	assert.Equal(t, got , nil )
}

func TestTaskDelete(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_port.NewMockTaskPort(mock)
	mockResult.EXPECT().Delete(1).AnyTimes().Return(nil)
	service := usecase.ProviderTaskDriver(mockResult)
	got := service.Delete(1)
	assert.Equal(t, got , nil )
}