package application

import "clockify/project/domain"

func (p *ProjectServiceImpl) GetByUserID(userID int) ([]domain.Projects, error) {
	return p.projectRepository.GetByUserID(userID)
}
