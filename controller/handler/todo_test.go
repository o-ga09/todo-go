package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-go/controller/handler"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
)

func TestGetTask(t *testing.T) {
	// テスト用のGinコンテキストを作成
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := &handler.TaskHandler{}
	router.GET("api/vi/task", handler.GetTask)

	// テスト用のリクエストを作成
	req, _ := http.NewRequest(http.MethodGet, "/api/vi/task", nil)
	w := httptest.NewRecorder()

	// テスト用のリクエストを実行
	router.ServeHTTP(w, req)

	// レスポンスのJSONデータをパースして検証
	var response ResponseJSON
	json.Unmarshal(w.Body.Bytes(), &response)
	
	// TODO: 期待値をユーズケースの結果に変更する
	expected := ResponseJSON{Taskid: 1}
	assert.Equal(t,expected,response)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetTasks(t *testing.T) {
	// テスト用のGinコンテキストを作成
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := &handler.TaskHandler{}
	router.GET("api/vi/tasks", handler.GetTasks)

	// テスト用のリクエストを作成
	req, _ := http.NewRequest(http.MethodGet, "/api/vi/tasks", nil)
	w := httptest.NewRecorder()

	// テスト用のリクエストを実行
	router.ServeHTTP(w, req)

	// レスポンスのJSONデータをパースして検証
	var response ResponseJSON
	json.Unmarshal(w.Body.Bytes(), &response)
	
	// TODO: 期待値をユーズケースの結果に変更する
	expected := ResponseJSON{Taskid: 1}
	assert.Equal(t,expected,response)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRegisterTasks(t *testing.T) {
	// テスト用のGinコンテキストを作成
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := &handler.TaskHandler{}
	router.POST("api/vi/user", handler.CreateTask)

	// テスト用のリクエストを作成
	req, _ := http.NewRequest(http.MethodPost, "/api/vi/user", nil)
	w := httptest.NewRecorder()

	// テスト用のリクエストを実行
	router.ServeHTTP(w, req)

	// レスポンスのJSONデータをパースして検証
	var response ResponseJSON
	json.Unmarshal(w.Body.Bytes(), &response)
	
	// TODO: 期待値をユーズケースの結果に変更する
	expected := ResponseJSON{Taskid: 1}
	assert.Equal(t,expected,response)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateTask(t *testing.T) {
	// テスト用のGinコンテキストを作成
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := &handler.TaskHandler{}
	router.PATCH("api/vi/user/:id", handler.UpdateTask)

	// テスト用のリクエストを作成
	req, _ := http.NewRequest(http.MethodPatch, "/api/vi/user/1", nil)
	w := httptest.NewRecorder()

	// テスト用のリクエストを実行
	router.ServeHTTP(w, req)

	// レスポンスのJSONデータをパースして検証
	var response ResponseJSON
	json.Unmarshal(w.Body.Bytes(), &response)
	
	// TODO: 期待値をユーズケースの結果に変更する
	expected := ResponseJSON{Taskid: 1}
	assert.Equal(t,expected,response)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteTask(t *testing.T) {
	// テスト用のGinコンテキストを作成
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := &handler.TaskHandler{}
	router.DELETE("api/vi/user/:id", handler.DeleteTask)

	// テスト用のリクエストを作成
	req, _ := http.NewRequest(http.MethodDelete, "/api/vi/user/1", nil)
	w := httptest.NewRecorder()

	// テスト用のリクエストを実行
	router.ServeHTTP(w, req)

	// レスポンスのJSONデータをパースして検証
	var response ResponseJSON
	json.Unmarshal(w.Body.Bytes(), &response)
	
	// TODO: 期待値をユーズケースの結果に変更する
	expected := ResponseJSON{Taskid: 1}
	assert.Equal(t,expected,response)
	assert.Equal(t, http.StatusOK, w.Code)
}

type ResponseJSON struct {
	Taskid int
}