package adapter

import (
	"clockify/project/domain"
	"clockify/project/presentation/model"
)

func DomainToProject(project domain.Projects) model.Projects {
	return model.Projects{
		ID:      int(project.ID),
		Name:    project.Name,
		Client:  project.Client,
		Amount:  project.Amount,
		Tracked: project.Tracked,
	}
}
