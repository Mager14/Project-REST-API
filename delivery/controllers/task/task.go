package task

import (
	"Project-REST-API/delivery/controllers/common"
	"Project-REST-API/entities"
	"Project-REST-API/repository/task"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TaskController struct {
	repo task.Task
}

func New(repository task.Task) *TaskController {
	return &TaskController{
		repo: repository,
	}
}

func (tc *TaskController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := tc.repo.Get()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get All User", res))
	}
}

func (tc *TaskController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		taskId, _ := strconv.Atoi(c.Param("id"))

		res, err := tc.repo.GetById(taskId)

		if err != nil {
			return c.JSON(http.StatusNotFound, common.NotFound(http.StatusNotFound, "not found", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get User", res))
	}
}

func (tc *TaskController) TaskRegister() echo.HandlerFunc {
	return func(c echo.Context) error {
		task := RegisterRequestFormat{}

		if err := c.Bind(&task); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		res, err := tc.repo.TaskRegister(entities.Task{Nama: task.Nama})

		if err != nil {
			return c.JSON(http.StatusNotFound, common.InternalServerError())
		}

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success Create User", res))
	}
}

func (uc *TaskController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newUser = UpdateRequestFormat{}
		taskId, _ := strconv.Atoi(c.Param("id"))

		if err := c.Bind(&newUser); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		res, err := uc.repo.Update(taskId, entities.Task{Nama: newUser.Nama})

		if err != nil {
			return c.JSON(http.StatusNotFound, common.InternalServerError())
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Update User", res))
	}
}

func (uc *TaskController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		taskId, _ := strconv.Atoi(c.Param("id"))

		err := uc.repo.Delete(taskId)

		if err != nil {
			return c.JSON(http.StatusNotFound, common.InternalServerError())
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Delete User", nil))
	}
}
