package user

import (
	"Project-REST-API/entities"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {

	t.Run("UserGet", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(MockUserRepository{})
		userController.Get()(context)

		var response GetUsersResponseFormat

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, response.Data[0].Nama, "Adlan")
		//
	})
	t.Run("GetUser", func(t *testing.T) {
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

	// e := echo.New()
	// req := httptest.NewRequest(http.MethodGet, "/", nil)
	// res := httptest.NewRecorder()
	// context := e.NewContext(req, res)
	// context.SetPath("/users")

	// userController := New(&MockUserRepository{})
	// userController.Get()(context)

	// response := GetUserResponseFormat{}

	// json.Unmarshal([]byte(res.Body.Bytes()), &response)

	// assert.Equal(t, 200, response.Code)
	// assert.Equal(t, "Adlan", response.Data[0].Nama)

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
	t.Run("Eror GetById", func(t *testing.T) {
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
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockUserRepository{})
		userController.UserRegister()(context)

		response := RegisterUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 201, response.Code)
		assert.Equal(t, "Adlan", response.Data.Nama)
		assert.Equal(t, "adlan@adlan.com", response.Data.Email)
		assert.Equal(t, "adlan123", response.Data.Password)

	})
	t.Run("Eror UserRegister", func(t *testing.T) {
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

}
func TestUpdate(t *testing.T) {
	t.Run("Update", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		userController := New(&MockUserRepository{})
		userController.Update()(context)

		response := UpdateResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "Adlan", response.Data.Nama)

	})
}
func TestDelete(t *testing.T) {
	t.Run("UserRegister", func(t *testing.T) {
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
}

type MockUserRepository struct{}

func (m MockUserRepository) Get() ([]entities.User, error) {
	return []entities.User{
		{Nama: "Adlan", Email: "adlan@adlan.com", Password: "adlan123"},
	}, nil
}

func (m MockUserRepository) GetById(userId int) (entities.User, error) {
	return entities.User{Nama: "Adlan", Email: "adlan@adlan.com", Password: "adlan123"}, nil
}

func (m MockUserRepository) UserRegister(newUser entities.User) (entities.User, error) {
	return entities.User{Nama: "Adlan", Email: "adlan@adlan.com", Password: "adlan123"}, nil
}

func (m MockUserRepository) Login(data entities.User) (entities.User, error) {
	return entities.User{Email: "adlan@adlan.com", Password: "adlan123"}, nil
}

func (m MockUserRepository) Update(userId int, newUser entities.User) (entities.User, error) {
	return entities.User{Nama: "Adlan", Email: "adlan@adlan.com", Password: "adlan123"}, nil
}

func (m MockUserRepository) Delete(userId int) error {
	return nil
}

type MockFalseUserRepository struct{}

func (m MockFalseUserRepository) Get() ([]entities.User, error) {
	return nil, errors.New("False User Object")
}
func (m MockFalseUserRepository) GetById(userId int) (entities.User, error) {
	return entities.User{}, errors.New("False Get Object")
}
func (m MockFalseUserRepository) UserRegister(newUser entities.User) (entities.User, error) {
	return entities.User{}, errors.New("False Register Object")
}
func (m MockFalseUserRepository) Login(data entities.User) (entities.User, error) {
	return entities.User{}, errors.New("False Login Object")
}
func (m MockFalseUserRepository) Update(userId int, newUser entities.User) (entities.User, error) {
	return entities.User{}, errors.New("False Update Object")
}
func (m MockFalseUserRepository) Delete(userId int) error {
	return errors.New("False Delete Object")
}
