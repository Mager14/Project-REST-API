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

	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	repo := New(db)

	t.Run("Success Create User", func(t *testing.T) {
		mockUser := entities.User{Nama: "Steven", Email: "steven@steven.com", Password: "steven123"}
		res, err := repo.UserRegister(mockUser)
		assert.Nil(t, err)
		assert.Equal(t, 1, int(res.ID))
	})

	t.Run("Fail Create User", func(t *testing.T) {
		mockUser := entities.User{Model: gorm.Model{ID: 1}, Nama: "Steven", Email: "steven@steven.com", Password: "steven123"}
		_, err := repo.UserRegister(mockUser)
		assert.NotNil(t, err)
	})
}

func TestGet(t *testing.T) {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	repo := New(db)
	mockUser := entities.User{Nama: "Steven", Email: "steven@steven.com", Password: "steven123"}
	db.Create(&mockUser)

	t.Run("Success Getting User", func(t *testing.T) {
		res, err := repo.Get()
		assert.Nil(t, err)
		assert.Equal(t, mockUser.Nama, res[0].Nama)
	})

	db.Migrator().DropTable(&entities.User{})

	t.Run("Fail Getting User", func(t *testing.T) {
		_, err := repo.Get()
		assert.NotNil(t, err)
	})
}

func TestGetById(t *testing.T) {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	repo := New(db)
	mockUser := entities.User{Nama: "Steven", Email: "steven@steven.com", Password: "steven123"}
	db.Create(&mockUser)

	t.Run("Success Getting User ID", func(t *testing.T) {
		res, err := repo.GetById(int(mockUser.ID))
		assert.Nil(t, err)
		assert.Equal(t, mockUser.Nama, res.Nama)
	})

	db.Migrator().DropTable(&entities.User{})

	t.Run("Fail Getting User ID", func(t *testing.T) {
		_, err := repo.GetById(int(mockUser.ID))
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

	t.Run("Fail Getting Task ID", func(t *testing.T) {
		err := repo.Delete(int(mockTask.ID))
		assert.NotNil(t, err)
	})
}
