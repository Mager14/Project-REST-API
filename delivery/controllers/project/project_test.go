package project

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

	t.Run("GET", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/projects")

		projectController := New(MockProjectRepository{})
		projectController.Get()(context)

		var response GetProjectsResponseFormat

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "ProjectKu", response.Data[0].Nama)
		//
	})
	t.Run("GET", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/projects")

		falseprojectController := New(MockFalseProjectRepository{})
		falseprojectController.Get()(context)

		var response GetProjectsResponseFormat

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
		context.SetPath("/projects/:id")

		projectController := New(&MockProjectRepository{})
		projectController.GetById()(context)

		response := GetProjectResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "ProjectKu", response.Data.Nama)

	})
	t.Run("ErorGetById", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/projects/:id")

		falseprojectController := New(MockFalseProjectRepository{})
		falseprojectController.GetById()(context)

		var response GetProjectResponseFormat

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, response.Message, "not found")
	})

}

func TestProjectRegister(t *testing.T) {
	t.Run("ProjectRegister", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama": "ProjectKu",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/projects/register")

		projectController := New(MockProjectRepository{})
		projectController.ProjectRegister()(context)

		response := RegisterProjectResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		// assert.Equal(t, 201, response.Code)
		assert.Equal(t, "ProjectKu", response.Data.Nama)

	})
	t.Run("ErorProjectRegister", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/projects/register")

		projectController := New(MockFalseProjectRepository{})
		projectController.ProjectRegister()(context)

		response := RegisterProjectResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "There is some error on server", response.Message)

	})

	t.Run("ProjectRegisterBind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		fmt.Println(req)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/projects/register")

		projectController := New(MockFalseProjectRepository{})
		projectController.ProjectRegister()(context)

		response := RegisterProjectResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 400, response.Code)

	})

}

func TestUpdate(t *testing.T) {
	t.Run("Update", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama": "ProjectKu",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))

		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/projects/:id")

		projectController := New(&MockProjectRepository{})
		projectController.Update()(context)

		response := UpdateResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "ProjectKu", response.Data.Nama)

	})

	t.Run("ErrorUpdate", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/projects/:id")

		projectController := New(&MockFalseProjectRepository{})
		projectController.Update()(context)

		response := UpdateResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "There is some error on server", response.Message)

	})
	t.Run("UpdateBind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))

		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/projects/:id")

		tastController := New(&MockFalseProjectRepository{})
		tastController.Update()(context)

		response := UpdateResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 400, response.Code)

	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/projects/:id")

		projectController := New(&MockProjectRepository{})
		projectController.Delete()(context)

		response := DeleteResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, nil, response.Data)

	})

	t.Run("ErrorDelete", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/projects/:id")

		projectController := New(&MockFalseProjectRepository{})
		projectController.Delete()(context)

		response := DeleteResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "There is some error on server", response.Message)

	})
}

type MockProjectRepository struct{}

func (m MockProjectRepository) Get() ([]entities.Project, error) {
	return []entities.Project{
		{Nama: "ProjectKu"},
	}, nil
}

func (m MockProjectRepository) GetById(project_id int) (entities.Project, error) {
	return entities.Project{Nama: "ProjectKu"}, nil
}

func (m MockProjectRepository) ProjectRegister(newProject entities.Project) (entities.Project, error) {
	return entities.Project{Nama: "ProjectKu"}, nil
}

func (m MockProjectRepository) Update(project_id int, newProject entities.Project) (entities.Project, error) {
	return entities.Project{Nama: "ProjectKu"}, nil
}

func (m MockProjectRepository) Delete(project_id int) error {
	return nil
}

type MockFalseProjectRepository struct{}

func (m MockFalseProjectRepository) Get() ([]entities.Project, error) {
	return nil, errors.New("False Project Object")
}
func (m MockFalseProjectRepository) GetById(project_id int) (entities.Project, error) {
	return entities.Project{}, errors.New("False Get Object")
}
func (m MockFalseProjectRepository) ProjectRegister(newProject entities.Project) (entities.Project, error) {
	return entities.Project{}, errors.New("False Register Object")
}
func (m MockFalseProjectRepository) Update(project_id int, newProject entities.Project) (entities.Project, error) {
	return entities.Project{}, errors.New("False Update Object")
}
func (m MockFalseProjectRepository) Delete(project_id int) error {
	return errors.New("False Delete Object")
}
