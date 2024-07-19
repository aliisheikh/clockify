package application

import "errors"

func (p *ProjectServiceImpl) Delete(userID, projectID int) error {
	project, err := p.projectRepository.GetProjectByID(userID, projectID)

	if err != nil {
		return err
	}
	if project == nil {
		return errors.New("project not found")
	}
	err = p.projectRepository.DELETE(userID, projectID)
	if err != nil {
		return err
	}
	return nil
}
