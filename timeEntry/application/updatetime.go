package application

import (
	"clockify/timeEntry/domain"
	"errors"
)

func (t *TimeEntryServiceImpl) Update(timeEntry domain.TimeEntry) error {
	// Check if no fields are provided for the update
	if timeEntry.Description == "" && timeEntry.StartTime.IsZero() && timeEntry.EndTime.IsZero() {
		return errors.New("at least one field is required for the update")
	}

	// Validate that StartTime is before EndTime
	if !timeEntry.StartTime.IsZero() && !timeEntry.EndTime.IsZero() {
		if timeEntry.StartTime.After(*timeEntry.EndTime) {
			return errors.New("StartTime must be before EndTime")
		}
	}

	// Retrieve existing time entry data
	timeEntryData, err := t.timeEntryRepository.GetTimeEntryByID(timeEntry.UserID, timeEntry.ID)
	if err != nil {
		return err
	}

	// Update fields if provided
	if timeEntry.Description != "" {
		timeEntryData.Description = timeEntry.Description
	}
	if !timeEntry.StartTime.IsZero() {
		timeEntryData.StartTime = timeEntry.StartTime
	}
	if !timeEntry.EndTime.IsZero() {
		timeEntryData.EndTime = timeEntry.EndTime
	}

	// Save updated time entry data
	if err := t.timeEntryRepository.Save(timeEntryData); err != nil {
		return err
	}

	return nil
}
