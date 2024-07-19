package application

import (
	"clockify/project/domain"
	domain2 "clockify/users/domain"
	"errors"
	"fmt"
)

type ProjectServiceImpl struct {
	projectRepository domain.ProjectRepository
	userRepository    domain2.UserRepo
	//validate          *validator.Validate
}

func NewProjectServiceImpl(projectRepository domain.ProjectRepository, userRepository domain2.UserRepo) *ProjectServiceImpl {
	return &ProjectServiceImpl{
		projectRepository: projectRepository,
		userRepository:    userRepository,
	}
}

func (p *ProjectServiceImpl) Create(project domain.Projects) (*domain.Projects, error) {

	fmt.Println(project)
	if project.Name == "" {
		return &project, errors.New("project Name is required")
	}
	if project.Amount == 0 {
		return &project, errors.New("project Amount is required")
	}
	if project.Client == "" {
		return &project, errors.New("project Client is required")
	}
	if project.Tracked == 0 {
		return &project, errors.New("project Tracked is required")
	}

	projects := &domain.Projects{
		Name:    project.Name,
		Amount:  project.Amount,
		Client:  project.Client,
		UserID:  project.UserID,
		Tracked: project.Tracked,
	}
	userID := project.UserID
	projectID, err := p.projectRepository.SaveCreate(projects, *userID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	project.ID = projectID.ID

	return projects, nil

}
