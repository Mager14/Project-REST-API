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

		taskController := New(MockTaskRepository{})
		taskController.Get()(context)

		var response GetTasksResponseFormat

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

		falsetaskController := New(MockFalseTaskRepository{})
		falsetaskController.Get()(context)

		var response GetTaskResponseFormat

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

		taskController := New(&MockTaskRepository{})
		taskController.GetById()(context)

		response := GetTaskResponseFormat{}

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

		falsetaskController := New(MockFalseTaskRepository{})
		falsetaskController.GetById()(context)

		var response GetTaskResponseFormat

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, response.Message, "not found")
	})

}

func TestTaskRegister(t *testing.T) {
	t.Run("TaskRegister", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama":       "Mawan",
			"priority":   1,
			"user_ID":    1,
			"project_ID": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/task/register")

		taskController := New(MockTaskRepository{})
		taskController.TaskRegister()(context)

		response := RegisterTaskResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		// assert.Equal(t, 201, response.Code)
		assert.Equal(t, "Mawan", response.Data.Nama)

	})
	t.Run("ErorTaskRegister", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/tasks/register")

		userController := New(MockFalseTaskRepository{})
		userController.TaskRegister()(context)

		response := RegisterTaskResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "There is some error on server", response.Message)

	})

	t.Run("TaskRegisterBind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama":       "Mawan",
			"priority":   1,
			"user_ID":    1,
			"project_ID": "test",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		fmt.Println(req)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/tasks/register")

		userController := New(MockFalseTaskRepository{})
		userController.TaskRegister()(context)

		response := RegisterTaskResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 400, response.Code)

	})

}

func TestUpdate(t *testing.T) {
	t.Run("Update", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama":     "Adlan",
			"email":    "adlan@adlan.com",
			"password": "adlan123",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))

		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		userController := New(&MockUserRepository{})
		userController.Update()(context)

		response := UpdateResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "Adlan", response.Data.Nama)

	})

	t.Run("ErrorUpdate", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		userController := New(&MockFalseUserRepository{})
		userController.Update()(context)

		response := UpdateResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "There is some error on server", response.Message)

	})
	t.Run("UpdateBind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama":     "Adlan",
			"email":    "adlan@adlan.com",
			"password": 123,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))

		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		userController := New(&MockUserRepository{})
		userController.Update()(context)

		response := UpdateResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 400, response.Code)

	})
}

func TestDelete(t *testing.T) {
	t.Run("DeleteTask", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/tasks/:id")

		taskController := New(&MockTaskRepository{})
		taskController.Delete()(context)

		response := DeleteResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, nil, response.Data)

	})

	t.Run("ErrorDeleteTask", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/tasks/:id")

		userController := New(&MockFalseTaskRepository{})
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
		entities.Task{Nama: "Adlan", Priority: 1, User_ID: 1, Project_ID: 1},
	}, nil
}

func (m MockTaskRepository) GetById(taskId int) (entities.Task, error) {
	return entities.Task{Nama: "Adlan", Priority: 1, User_ID: 1, Project_ID: 1}, nil
}

func (m MockTaskRepository) TaskRegister(newTask entities.Task) (entities.Task, error) {
	return entities.Task{Nama: "Adlan", Priority: 1, User_ID: 1, Project_ID: 1}, nil
}

func (m MockTaskRepository) Update(taskId int, newTask entities.Task) (entities.Task, error) {
	return entities.Task{Nama: "Adlan", Priority: 1, User_ID: 1, Project_ID: 1}, nil
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
