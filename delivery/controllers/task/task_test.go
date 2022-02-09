package task

import (
	"Project-REST-API/entities"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {

	t.Run("ErrorGetUser", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(MockUserRepository{})
		userController.Get()(context)

		var response GetUsersResponseFormat

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Adlan", response.Data[0].Nama)
		//
	})
	t.Run("ErrorGetUser", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/users")

		falseUserController := New(MockFalseUserRepository{})
		falseUserController.Get()(context)

		var response GetUserResponseFormat

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, response.Message, "There is some error on server")
	})

}

func TestGetById(t *testing.T) {
	t.Run("GetById", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		userController := New(&MockUserRepository{})
		userController.GetById()(context)

		response := GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "Adlan", response.Data.Nama)

	})
	t.Run("ErorGetById", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		falseUserController := New(MockFalseUserRepository{})
		falseUserController.GetById()(context)

		var response GetUserResponseFormat

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, response.Message, "not found")
	})

}

func TestUserRegister(t *testing.T) {
	t.Run("UserRegister", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "Adlan",
			"email":    "adlan@adlan.com",
			"password": "adlan123",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(MockUserRepository{})
		userController.UserRegister()(context)

		response := RegisterUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		// assert.Equal(t, 201, response.Code)
		assert.Equal(t, "Adlan", response.Data.Nama)

	})
	t.Run("ErorUserRegister", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockFalseUserRepository{})
		userController.UserRegister()(context)

		response := RegisterUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "There is some error on server", response.Message)

	})

	t.Run("UserRegisterBind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama":     "Adlan",
			"email":    "adlan@adlan.com",
			"password": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		fmt.Println(req)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(MockUserRepository{})
		userController.UserRegister()(context)

		response := RegisterUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 400, response.Code)

	})

}

func TestUpdate(t *testing.T) {
	t.Run("Update Task", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama":       "Taskku",
			"priority":   12,
			"user_id":    1,
			"project_id": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))

		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/task/:id")

		taskController := New(&MockTaskRepository{})
		taskController.Update()(context)

		response := UpdateResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "Taskku", response.Data.Nama)

	})

	t.Run("ErrorUpdateTask", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/task/:id")

		taskController := New(&MockFalseTaskRepository{})
		taskController.Update()(context)

		response := UpdateResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "There is some error on server", response.Message)

	})
	t.Run("UpdateBind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama":       "Taskku",
			"priority":   12,
			"user_id":    1,
			"project_id": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))

		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/task/:id")

		userController := New(&MockTaskRepository{})
		userController.Update()(context)

		response := UpdateResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 400, response.Code)

	})
}

func TestDelete(t *testing.T) {
	t.Run("DeleteUser", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		userController := New(&MockUserRepository{})
		userController.Delete()(context)

		response := DeleteResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, nil, response.Data)

	})

	t.Run("ErrorDeleteUser", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		userController := New(&MockFalseUserRepository{})
		userController.Delete()(context)

		response := DeleteResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "There is some error on server", response.Message)

	})
}

type MockTaskRepository struct{}

func (m MockTaskRepository) Get() ([]entities.Task, error) {
	return []entities.Task{
		{Nama: "Taskku", Priority: 12, User_ID: 1, Project_ID: 1},
	}, nil
}

func (m MockTaskRepository) GetById(taskId int) (entities.Task, error) {
	return entities.Task{Nama: "Taskku", Priority: 12, User_ID: 1, Project_ID: 1}, nil
}

func (m MockTaskRepository) TaskRegister(newTask entities.Task) (entities.Task, error) {
	return entities.Task{Nama: "Taskku", Priority: 12, User_ID: 1, Project_ID: 1}, nil
}

func (m MockTaskRepository) Update(taskId int, newTask entities.Task) (entities.Task, error) {
	return entities.Task{Nama: "Taskku", Priority: 12, User_ID: 1, Project_ID: 1}, nil
}

func (m MockTaskRepository) Delete(taskId int) error {
	return nil
}

type MockFalseTaskRepository struct{}

func (m MockFalseTaskRepository) Get() ([]entities.Task, error) {
	return nil, errors.New("False Task Object")
}
func (m MockFalseTaskRepository) GetById(taskId int) (entities.Task, error) {
	return entities.Task{}, errors.New("False Get Object")
}
func (m MockFalseTaskRepository) TaskRegister(newTask entities.Task) (entities.Task, error) {
	return entities.Task{}, errors.New("False Register Object")
}
func (m MockFalseTaskRepository) Update(taskId int, newTask entities.Task) (entities.Task, error) {
	return entities.Task{}, errors.New("False Update Object")
}
func (m MockFalseTaskRepository) Delete(taskId int) error {
	return errors.New("False Delete Object")
}
