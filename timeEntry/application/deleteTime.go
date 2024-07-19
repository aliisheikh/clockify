package application

import "errors"

func (t *TimeEntryServiceImpl) Delete(userID, timeEntryID int) error {
	// Retrieve time entry by ID
	timeEntry, err := t.timeEntryRepository.GetTimeEntryByID(userID, timeEntryID)
	if err != nil {
		return err
	}

	// Check if time entry exists
	if timeEntry == nil {
		return errors.New("time entry not found")
	}

	// Delete time entry
	err = t.timeEntryRepository.DELETE(userID, timeEntryID)
	if err != nil {
		return err
	}

	return nil
}
