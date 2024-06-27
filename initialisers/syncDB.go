package initialisers

import (
	"github.com/awhb/gitgood-backend/models"
)

func SyncDB() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Thread{})
	DB.AutoMigrate(&models.Comment{})
}