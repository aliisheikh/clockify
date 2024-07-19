package domain

import (
	"time"
)

type TimeEntry struct {
	ID          int       `json:"id"`
	ProjectID   int       `json:"projectID"`
	UserID      int       `json:"userID"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	Description string    `json:"description"`
}

type TimeEntryRepository interface {
	DELETE(userID, timeID int) error
	Create(timeEntry *TimeEntry) error
	GetTimeEntryByID(userID, timeID int) (*TimeEntry, error)
	Update(userID int, ID int, updatedTimeEntry *TimeEntry) error
	Save(timeEntry *TimeEntry) error
	SaveCreate(timeEntry *TimeEntry, userID int) (*TimeEntry, error)
	GetByUserID(userID int) ([]TimeEntry, error)
}
