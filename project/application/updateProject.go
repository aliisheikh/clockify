package application

import (
	"clockify/project/domain"
	"errors"
)

func (p *ProjectServiceImpl) Update(project domain.Projects) error {
	// Check if no fields are provided for the update
	if project.Name == "" && project.Tracked == 0 && project.Amount == 0 && project.Client == "" {
		return errors.New("at least one field is required for the update")
	}

	// Validate that StartTime is before EndTime

	// Retrieve existing time entry data
	projectData, err := p.projectRepository.GetProjectByID(*project.UserID, project.ID)
	if err != nil {
		return err
	}

	// Update fields if provided
	if project.Name != "" {
		projectData.Name = project.Name
	}
	if project.Amount != 0 {
		projectData.Amount = project.Amount
	}
	if project.Client != "" {
		projectData.Client = project.Client
	}
	if project.Tracked != 0 {
		projectData.Tracked = project.Tracked
	}

	// Save updated time entry data
	if err := p.projectRepository.Save(projectData); err != nil {
		return err
	}

	return nil
}
