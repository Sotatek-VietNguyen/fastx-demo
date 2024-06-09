package seeders

import (
	"fastx-api/src/models"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) error {
	users := []models.User{
		{Name: "John Doe", Email: "john@example.com", Password: "password"},
		{Name: "Jane Doe", Email: "jane@example.com", Password: "password"},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}

	return nil
}
