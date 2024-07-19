package application

import (
	"clockify/timeEntry/domain"
	domain2 "clockify/users/domain"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type TimeEntryServiceImpl struct {
	timeEntryRepository domain.TimeEntryRepository
	userRepository      domain2.UserRepo
	validate            *validator.Validate
}

func NewTimeEntryServiceImpl(timeEntryRepository domain.TimeEntryRepository, userRepository domain2.UserRepo) *TimeEntryServiceImpl {
	return &TimeEntryServiceImpl{
		timeEntryRepository: timeEntryRepository,
		userRepository:      userRepository,
	}
}

func (t *TimeEntryServiceImpl) Create(timeEntry domain.TimeEntry) (*domain.TimeEntry, error) {
	// Validate required fields
	if timeEntry.Description == "" {
		return &timeEntry, errors.New("description is required")
	}
	if timeEntry.StartTime.IsZero() {
		return &timeEntry, errors.New("StartTime is required")
	}
	if timeEntry.EndTime.IsZero() {
		return &timeEntry, errors.New("EndTime is required")
	}
	if timeEntry.UserID == 0 {
		return &timeEntry, errors.New("UserID is required")
	}
	if timeEntry.ProjectID == 0 {
		return &timeEntry, errors.New("ProjectID is required")
	}

	// Create TimeEntry object
	newTimeEntry := &domain.TimeEntry{
		Description: timeEntry.Description,
		StartTime:   timeEntry.StartTime,
		EndTime:     timeEntry.EndTime,
		UserID:      timeEntry.UserID,
		ProjectID:   timeEntry.ProjectID,
	}

	// Call repository to save TimeEntry
	timeID, err := t.timeEntryRepository.SaveCreate(newTimeEntry, timeEntry.UserID)
	if err != nil {
		fmt.Println(err)
		return timeID, err
	}
	timeEntry.ID = timeID.ID

	return timeID, nil
}
