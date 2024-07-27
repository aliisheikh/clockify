package application

import (
	"clockify/timeEntry/domain"
	"time"
)

func (t *TimeEntryServiceImpl) StartTimeEntry(userID uint, description string) (*domain.TimeEntry, error) {
	timeEntry := &domain.TimeEntry{
		UserID:      int(userID),
		StartTime:   time.Now(),
		Description: description,
		ProjectID:   nil, // Initially no project
		EndTime:     nil,
	}

	err := t.timeEntryRepository.Save(timeEntry)
	if err != nil {
		return nil, err
	}

	return timeEntry, nil
}
