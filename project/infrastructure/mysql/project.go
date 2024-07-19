package mysql

import (
	"clockify/project/domain"
	domain2 "clockify/users/domain"
	"errors"
	"fmt"
	"gorm.io/gorm"
	//"clockify/users/domain"
)

type ProjectRepoImpl struct {
	DB *gorm.DB
}

func NewProjectRepoImpl(db *gorm.DB) *ProjectRepoImpl {
	return &ProjectRepoImpl{DB: db}
}

func (p *ProjectRepoImpl) DELETE(userID, projID int) error {
	project := domain.Projects{}
	result := p.DB.Where("ID = ? AND user_id = ?", projID, userID).First(&project)
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

func (p *ProjectRepoImpl) GetProjectByID(userID, projID int) (*domain.Projects, error) {
	if userID == 0 {
		return nil, errors.New("userID is required")
	}
	if projID == 0 {
		return nil, errors.New("projID is required")
	}
	var projects domain.Projects
	result := p.DB.First(&projects, "ID=? AND user_id=?", projID, userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("project not found")
		}
		return nil, result.Error
	}
	return &projects, nil
}

func (p *ProjectRepoImpl) Save(project *domain.Projects) error {
	existingProject, err := p.GetProjectByID(*project.UserID, project.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := p.DB.Create(&project)
			if result.Error != nil {
				return result.Error
			}
			return nil
		}
		return err
	}
	result := p.DB.Model(&existingProject).Save(&project)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (p *ProjectRepoImpl) SaveCreate(project *domain.Projects, userID int) (*domain.Projects, error) {
	// Set the UserID of the project
	project.UserID = &userID

	// Check if the project already exists
	var existingProject domain.Projects
	result := p.DB.Where("id = ?", project.ID).First(&existingProject)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return project, fmt.Errorf("failed to retrieve project: %w", result.Error)
		}
	} else {
		// If a project with the same ID exists, return an error
		return project, errors.New("project with the given ID already exists")
	}

	// Create the new project
	result = p.DB.Create(project)
	if result.Error != nil {
		return project, fmt.Errorf("failed to save project: %w", result.Error)
	}

	return project, nil
}

func (p *ProjectRepoImpl) GetByUserID(userID int) ([]domain.Projects, error) {
	var projects []domain.Projects
	var user domain2.User

	if err := p.DB.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with ID %d not found", userID)
		}
		return nil, fmt.Errorf("failed to retrieve the user: %w", err)
	}
	if err := p.DB.Model(&domain.Projects{}).Where("user_id = ?", userID).Find(&projects).Error; err != nil {
		return nil, err
	}
	//for i := range projects {
	//	projects[i].UserID = projects[i].UserID
	//}
	return projects, nil
}

func (p *ProjectRepoImpl) Update(userID int, ID int, updatedProject *domain.Projects) error {
	// Check if the TimeEntry exists for the given userID
	var existingProject domain.Projects
	if err := p.DB.Where("id = ? AND user_id = ?", ID, userID).First(&existingProject).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("time entry with ID %s not found for user %d", ID, userID)
		}
		return err
	}

	// Update the fields of existingTimeEntry with the new values
	//existingTimeEntry.ProjectID = updatedTimeEntry.ProjectID
	existingProject.UserID = updatedProject.UserID
	existingProject.Name = updatedProject.Name
	existingProject.Client = updatedProject.Client
	existingProject.Amount = updatedProject.Amount
	existingProject.Tracked = updatedProject.Tracked

	// Save the updated TimeEntry back to the database
	if err := p.DB.Save(&existingProject).Error; err != nil {
		return err
	}

	return nil
}
