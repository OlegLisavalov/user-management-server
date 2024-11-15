package services

import (
	"errors"
	"user-management-server/database"
	"user-management-server/models"
)

func CompleteTask(userID int, taskName string) error {
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return errors.New("user not found")
	}

	task := models.Task{Name: taskName, Points: 10, UserID: uint(userID), Completed: true}
	user.Points += task.Points
	if err := database.DB.Save(&user).Error; err != nil {
		return errors.New("failed to update user points")
	}
	return database.DB.Create(&task).Error
}
