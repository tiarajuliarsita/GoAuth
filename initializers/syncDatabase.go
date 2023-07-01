package initializers

import "lear-jwt/models"

func SyncDatabase() {
	users := models.User{}
	DB.AutoMigrate(&users)
}
