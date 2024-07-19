package domain

type ProjectService interface {
	Create(project Projects) (*Projects, error)
	Delete(userID, projectID int) error
	GetByUserID(userID int) ([]Projects, error)
	GetProjectByID(userID, projectID int) (*Projects, error)
	Update(timeEntry Projects) error
}
