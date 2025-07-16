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
