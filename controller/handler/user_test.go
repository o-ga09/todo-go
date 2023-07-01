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

func TestGetUser(t *testing.T) {
	// テスト用のGinコンテキストを作成
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := &handler.UserHnadler{}
	router.GET("api/vi/user", handler.GetUser)

	// テスト用のリクエストを作成
	req, _ := http.NewRequest(http.MethodGet, "/api/vi/user", nil)
	w := httptest.NewRecorder()

	// テスト用のリクエストを実行
	router.ServeHTTP(w, req)

	// レスポンスのJSONデータをパースして検証
	var response ResponseUser
	json.Unmarshal(w.Body.Bytes(), &response)
	
	// TODO: 期待値をユーズケースの結果に変更する
	expected := ResponseUser{Userid: 1}
	assert.Equal(t,expected,response)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetUsers(t *testing.T) {
	// テスト用のGinコンテキストを作成
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := &handler.UserHnadler{}
	router.GET("api/vi/user/:id", handler.GetUsers)

	// テスト用のリクエストを作成
	req, _ := http.NewRequest(http.MethodGet, "/api/vi/user/1", nil)
	w := httptest.NewRecorder()

	// テスト用のリクエストを実行
	router.ServeHTTP(w, req)

	// レスポンスのJSONデータをパースして検証
	var response ResponseUser
	json.Unmarshal(w.Body.Bytes(), &response)
	
	// TODO: 期待値をユーズケースの結果に変更する
	expected := ResponseUser{Userid: 1}
	assert.Equal(t,expected,response)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRegisterUser(t *testing.T) {
	// テスト用のGinコンテキストを作成
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := &handler.UserHnadler{}
	router.POST("api/vi/user", handler.CreateUser)

	// テスト用のリクエストを作成
	req, _ := http.NewRequest(http.MethodPost, "/api/vi/user", nil)
	w := httptest.NewRecorder()

	// テスト用のリクエストを実行
	router.ServeHTTP(w, req)

	// レスポンスのJSONデータをパースして検証
	var response ResponseUser
	json.Unmarshal(w.Body.Bytes(), &response)
	
	// TODO: 期待値をユーズケースの結果に変更する
	expected := ResponseUser{Userid: 1}
	assert.Equal(t,expected,response)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateUser(t *testing.T) {
	// テスト用のGinコンテキストを作成
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := &handler.UserHnadler{}
	router.PATCH("api/vi/user/:id", handler.UpdateUser)

	// テスト用のリクエストを作成
	req, _ := http.NewRequest(http.MethodPatch, "/api/vi/user/1", nil)
	w := httptest.NewRecorder()

	// テスト用のリクエストを実行
	router.ServeHTTP(w, req)

	// レスポンスのJSONデータをパースして検証
	var response ResponseUser
	json.Unmarshal(w.Body.Bytes(), &response)
	
	// TODO: 期待値をユーズケースの結果に変更する
	expected := ResponseUser{Userid: 1}
	assert.Equal(t,expected,response)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteUser(t *testing.T){
	// テスト用のGinコンテキストを作成
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := &handler.UserHnadler{}
	router.DELETE("api/vi/user/:id", handler.DeleteUser)

	// テスト用のリクエストを作成
	req, _ := http.NewRequest(http.MethodDelete, "/api/vi/user/1", nil)
	w := httptest.NewRecorder()

	// テスト用のリクエストを実行
	router.ServeHTTP(w, req)

	// レスポンスのJSONデータをパースして検証
	var response ResponseUser
	json.Unmarshal(w.Body.Bytes(), &response)
	
	// TODO: 期待値をユーズケースの結果に変更する
	expected := ResponseUser{Userid: 1}
	assert.Equal(t,expected,response)
	assert.Equal(t, http.StatusOK, w.Code)
}

type ResponseUser struct {
	Userid int
}