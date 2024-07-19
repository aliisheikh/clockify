package domain

type Projects struct {
	ID      int
	Name    string  `json:"name"`
	Client  string  `json:"client"`
	Amount  float32 `json:"amount"`
	Tracked float32 `json:"tracked"`
	//CreatedAt time.Time
	UserID *int `json:"user_id"`
}

type ProjectRepository interface {
	DELETE(userID, projID int) error
	GetProjectByID(userID, projID int) (*Projects, error)
	Save(project *Projects) error
	SaveCreate(project *Projects, userID int) (*Projects, error)
	GetByUserID(userID int) ([]Projects, error)
	Update(userID int, ID int, updatedProject *Projects) error
}
