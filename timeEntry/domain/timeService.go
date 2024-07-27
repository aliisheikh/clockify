package domain

type TimeEntryService interface {
	Create(timeEntry TimeEntry) (*TimeEntry, error)
	Update(timeEntry TimeEntry) error
	Delete(userID, timeEntryID int) error
	GetByUserID(userID int) ([]TimeEntry, error)
	GetTimeEntryByID(userID, timeEntryID int) (*TimeEntry, error)
	StartTimeEntry(userID uint, description string) (*TimeEntry, error)
}
