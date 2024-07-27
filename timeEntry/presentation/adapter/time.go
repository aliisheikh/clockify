package adapter

import (
	"clockify/timeEntry/domain"
	"clockify/timeEntry/presentation/models"
)

func DomainToTime(time domain.TimeEntry) models.TimeEntry {
	return models.TimeEntry{
		ID:          time.ID,
		ProjectID:   *time.ProjectID,
		UserID:      time.UserID,
		StartTime:   time.StartTime,
		EndTime:     *time.EndTime,
		Description: time.Description,
	}
}
