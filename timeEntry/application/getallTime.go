package application

import "clockify/timeEntry/domain"

func (t *TimeEntryServiceImpl) GetByUserID(userID int) ([]domain.TimeEntry, error) {
	// Retrieve list of time entries from repository
	timeEntries, err := t.timeEntryRepository.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	// Iterate through each time entry to fetch associated user data
	for i := range timeEntries {
		userData, err := t.userRepository.GetUserByID(timeEntries[i].UserID)
		if err != nil {
			return nil, err
		}

		// Populate user-related fields in time entry object
		timeEntries[i].UserID = userData.ID // Assuming Name is a field in domain.TimeEntry to store user's name
		// You can populate other user-related fields as needed
	}

	return timeEntries, nil
}
