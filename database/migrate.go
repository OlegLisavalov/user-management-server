package database

import (
	"fmt"
	"log"
	"user-management-server/models"
)

func MigrateDB() error {
	if err := DB.AutoMigrate(&models.User{}, &models.Task{}); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
	log.Println("Database migration successful")
	return nil
}
