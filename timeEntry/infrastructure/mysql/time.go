package mysql

import (
	"clockify/timeEntry/domain"
	entity3 "clockify/timeEntry/infrastructure/entity"
	domain2 "clockify/users/domain"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type TimeEntryRepoImp struct {
	DB *gorm.DB
}

func NewTimeEntryRepoImp(db *gorm.DB) *TimeEntryRepoImp {
	return &TimeEntryRepoImp{DB: db}
}

func (p *TimeEntryRepoImp) DELETE(userID, timeID int) error {
	project := domain.TimeEntry{}
	result := p.DB.Where("ID = ? AND user_id = ?", timeID, userID).First(&project)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("project not found")
		}
		return result.Error
	}

	result = p.DB.Delete(&project)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *TimeEntryRepoImp) Create(timeEntry *domain.TimeEntry) error {
	if err := t.DB.Create(timeEntry).Error; err != nil {
		return err
	}
	return nil
}

func (p *TimeEntryRepoImp) GetTimeEntryByID(userID, timeID int) (*domain.TimeEntry, error) {
	if userID == 0 {
		return nil, errors.New("userID is required")
	}
	if timeID == 0 {
		return nil, errors.New("projID is required")
	}
	var timeEntry domain.TimeEntry
	result := p.DB.First(&timeEntry, "ID=? AND user_id=?", timeID, userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("timeEntry not found")
		}
		return nil, result.Error
	}
	return &timeEntry, nil
}

func (t *TimeEntryRepoImp) Update(userID int, ID int, updatedTimeEntry *domain.TimeEntry) error {
	// Check if the TimeEntry exists for the given userID
	var existingTimeEntry domain.TimeEntry
	if err := t.DB.Where("id = ? AND user_id = ?", ID, userID).First(&existingTimeEntry).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("time entry with ID %s not found for user %d", ID, userID)
		}
		return err
	}

	// Update the fields of existingTimeEntry with the new values
	existingTimeEntry.ProjectID = updatedTimeEntry.ProjectID
	existingTimeEntry.UserID = updatedTimeEntry.UserID
	existingTimeEntry.StartTime = updatedTimeEntry.StartTime
	existingTimeEntry.EndTime = updatedTimeEntry.EndTime
	existingTimeEntry.Description = updatedTimeEntry.Description

	// Save the updated TimeEntry back to the database
	if err := t.DB.Save(&existingTimeEntry).Error; err != nil {
		return err
	}

	return nil
}

func (t *TimeEntryRepoImp) Save(timeEntry *domain.TimeEntry) error {
	var existingTimeEntry domain.TimeEntry

	// Check if a time entry with the same ID already exists
	if err := t.DB.First(&existingTimeEntry, "id = ?", timeEntry.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// If no existing entry is found, create a new one
			if err := t.DB.Create(&timeEntry).Error; err != nil {
				return err
			}
			return nil
		}
		return err
	}

	// If an existing entry is found, update it
	if err := t.DB.Model(&existingTimeEntry).Updates(timeEntry).Error; err != nil {
		return err
	}
	return nil
}

func (t *TimeEntryRepoImp) SaveCreate(timeEntry *domain.TimeEntry, userID int) (*domain.TimeEntry, error) {
	// Set the UserID of the project
	timeEntry.UserID = userID

	// Check if the project already exists
	var existingProject entity3.TimeEntry
	result := t.DB.Where("id = ?", timeEntry.ID).First(&existingProject)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return timeEntry, fmt.Errorf("failed to retrieve timeEntry: %w", result.Error)
		}
	} else {
		// If a project with the same ID exists, return an error
		return timeEntry, errors.New("time Entry with the given ID already created")
	}

	// Create the new project
	result = t.DB.Create(timeEntry)
	if result.Error != nil {
		return timeEntry, fmt.Errorf("failed to save project: %w", result.Error)
	}

	return timeEntry, nil
}

func (p *TimeEntryRepoImp) GetByUserID(userID int) ([]domain.TimeEntry, error) {
	var projects []domain.TimeEntry
	var user domain2.User

	if err := p.DB.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with ID %d not found", userID)
		}
		return nil, fmt.Errorf("failed to retrieve the user: %w", err)
	}
	if err := p.DB.Model(&domain.TimeEntry{}).Where("user_id = ?", userID).Find(&projects).Error; err != nil {
		return nil, err
	}
	//for i := range projects {
	//	projects[i].UserID = projects[i].UserID
	//}
	return projects, nil
}
