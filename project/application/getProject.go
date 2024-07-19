package application

import "clockify/project/domain"

func (p *ProjectServiceImpl) GetProjectByID(userID, projectID int) (*domain.Projects, error) {
	// Retrieve project data from repository
	projectData, err := p.projectRepository.GetProjectByID(userID, projectID)
	if err != nil {
		return nil, err
	}

	// Retrieve user data associated with the project
	userData, err := p.userRepository.GetUserByID(*projectData.UserID)
	if err != nil {
		return nil, err
	}

	// Create domain.Project response with combined data
	projectResponse := &domain.Projects{
		ID:      projectData.ID,
		Name:    projectData.Name,
		Client:  projectData.Client,
		Amount:  projectData.Amount,
		Tracked: projectData.Tracked,
		//CreatedAt: projectData.CreatedAt,
		UserID: &userData.ID,
		// Assuming you have a Username field in domain.Project that you want to populate

		// Other fields from userData can be mapped as needed
	}

	return projectResponse, nil
}
