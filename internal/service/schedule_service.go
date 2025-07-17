package service

import (
	"fmt"
	"go-sql-project/internal/model"
)

func CreateSchedule(schedule *model.Schedule) (*model.Schedule, error) {
	if err := DB.Create(&schedule).Error; err != nil {
		return nil, fmt.Errorf("error when creating schedule: %v", err)
	}

	return schedule, nil
}

func EditSchedule(schedule *model.Schedule) (*model.Schedule, error) {
	var currentSchedule model.Schedule
	if err := DB.Where("id = ?", schedule.ID).First(&currentSchedule).Error; err != nil {
		return schedule, err
	}

	if err := DB.Model(&currentSchedule).Updates(schedule).Error; err != nil {
		return schedule, err
	}
	return schedule, nil
}
