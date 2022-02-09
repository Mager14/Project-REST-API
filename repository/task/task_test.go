package task

import (
	"Project-REST-API/configs"
	"Project-REST-API/entities"
	"Project-REST-API/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestInsert(t *testing.T) {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Task{})
	db.AutoMigrate(&entities.Task{})

	repo := New(db)

	t.Run("Success Creating Task", func(t *testing.T) {
		mockTask := entities.Task{Nama: "Steven"}
		res, err := repo.TaskRegister(mockTask)
		assert.Nil(t, err)
		assert.Equal(t, 1, int(res.ID))
	})

	t.Run("Fail Creating Task", func(t *testing.T) {
		mockTask := entities.Task{Model: gorm.Model{ID: 1}, Nama: "Steven", Priority: 1, User_ID: 1, Project_ID: 1}
		_, err := repo.TaskRegister(mockTask)
		assert.NotNil(t, err)
	})
}

func TestGet(t *testing.T) {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Task{})
	db.AutoMigrate(&entities.Task{})

	repo := New(db)
	mockTask := entities.Task{Nama: "Steven", Priority: 1, User_ID: 1, Project_ID: 1}
	db.Create(&mockTask)

	t.Run("Success Getting Task", func(t *testing.T) {
		res, err := repo.Get()
		assert.Nil(t, err)
		assert.Equal(t, mockTask.Nama, res[0].Nama)
	})

	db.Migrator().DropTable(&entities.Task{})

	t.Run("Fail Getting Task", func(t *testing.T) {
		_, err := repo.Get()
		assert.NotNil(t, err)
	})
}

func TestGetById(t *testing.T) {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Task{})
	db.AutoMigrate(&entities.Task{})

	repo := New(db)
	mockTask := entities.Task{Nama: "Mawan", Priority: 1, User_ID: 1, Project_ID: 1}
	db.Create(&mockTask)

	t.Run("Success Getting Task by ID", func(t *testing.T) {
		res, err := repo.GetById(int(mockTask.ID))
		assert.Nil(t, err)
		assert.Equal(t, mockTask.Nama, res.Nama)
	})

	db.Migrator().DropTable(&entities.Task{})

	t.Run("Fail Getting Task by ID", func(t *testing.T) {
		_, err := repo.GetById(int(mockTask.ID))
		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Task{})
	db.AutoMigrate(&entities.Task{})

	repo := New(db)
	mockTask := entities.Task{Nama: "Task1", Priority: 12}
	db.Create(&mockTask)

	t.Run("Success Update Task", func(t *testing.T) {
		mockUpdate := entities.Task{Nama: "Task1", Priority: 12}
		res, err := repo.Update(1, mockUpdate)
		assert.Nil(t, err)
		assert.Equal(t, mockUpdate.Nama, res.Nama)
	})

	t.Run("Fail Update Task", func(t *testing.T) {
		mockUpdate := entities.Task{Model: gorm.Model{ID: 1}, Nama: "Task1", Priority: 12}
		_, err := repo.Update(2, mockUpdate)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Task{})
	db.AutoMigrate(&entities.Task{})

	repo := New(db)
	mockTask := entities.Task{Nama: "Task1", Priority: 12}
	db.Create(&mockTask)

	t.Run("Success Deleting Task ID", func(t *testing.T) {
		err := repo.Delete(int(mockTask.ID))
		assert.Nil(t, err)
	})

	db.Migrator().DropTable(&entities.Task{})

	t.Run("Fail Deleting Task ID", func(t *testing.T) {
		err := repo.Delete(int(mockTask.ID))
		assert.NotNil(t, err)
	})
}
