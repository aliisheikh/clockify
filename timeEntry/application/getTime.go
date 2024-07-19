package application

import "clockify/timeEntry/domain"

func (t *TimeEntryServiceImpl) GetTimeEntryByID(userID, timeEntryID int) (*domain.TimeEntry, error) {
	// Retrieve time entry data from repository
	timeEntryData, err := t.timeEntryRepository.GetTimeEntryByID(userID, timeEntryID)
	if err != nil {
		return nil, err
	}

	// Retrieve user data associated with the time entry
	userData, err := t.userRepository.GetUserByID(timeEntryData.UserID)
	if err != nil {
		return nil, err
	}

	// Create domain.TimeEntry response with combined data
	timeEntryResponse := &domain.TimeEntry{
		ID:          timeEntryData.ID,
		Description: timeEntryData.Description,
		StartTime:   timeEntryData.StartTime,
		EndTime:     timeEntryData.EndTime,
		UserID:      userData.ID,
		ProjectID:   timeEntryData.ProjectID,
		// Assuming you have other fields in domain.TimeEntry that you want to populate

		// Other fields from userData can be mapped as needed
	}

	return timeEntryResponse, nil
}
