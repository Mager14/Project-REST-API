package task

import (
	"Project-REST-API/entities"

	"gorm.io/gorm"
)

type TaskRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		database: db,
	}
}

func (ur *TaskRepository) Get() ([]entities.Task, error) {
	arrTask := []entities.Task{}

	if err := ur.database.Find(&arrTask).Error; err != nil {
		return nil, err
	}

	return arrTask, nil
}

func (ur *TaskRepository) GetById(taskId int) (entities.Task, error) {
	arrTask := entities.Task{}

	if err := ur.database.First(&arrTask, taskId).Error; err != nil {
		return arrTask, err
	}

	return arrTask, nil
}

func (ur *TaskRepository) TaskRegister(u entities.Task) (entities.Task, error) {
	if err := ur.database.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (ur *TaskRepository) Update(taskId int, newTask entities.Task) (entities.Task, error) {

	var task entities.Task
	ur.database.First(&task, taskId)

	if err := ur.database.Model(&task).Updates(&newTask).Error; err != nil {
		return task, err
	}

	return task, nil
}

func (ur *TaskRepository) Delete(taskId int) error {

	var task entities.Task

	if err := ur.database.First(&task, taskId).Error; err != nil {
		return err
	}
	ur.database.Delete(&task, taskId)
	return nil

}
