package services

import (
	"errors"
	"user-management-server/database"
	"user-management-server/models"
)

func CompleteTask(userID string, taskType string) error {
	var user models.User
	if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
		return errors.New("user not found")
	}

	points := GetPointsForTask(taskType)

	task := models.Task{
		Name:      "Task for " + taskType,
		Type:      taskType,
		Points:    points,
		UserID:    userID,
		Completed: true,
	}

	user.Points += points
	if err := database.DB.Save(&user).Error; err != nil {
		return errors.New("failed to update user points")
	}

	return database.DB.Create(&task).Error
}

func GetPointsForTask(taskType string) int {
	switch taskType {
	case "subscribe in telegram":
		return 15
	case "subscribe in twitter":
		return 10
	case "subscribe in youtube":
		return 10
	case "like post in telegram":
		return 5
	default:
		return 5
	}
}
