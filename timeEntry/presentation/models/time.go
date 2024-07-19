package models

import "time"

type TimeEntry struct {
	ID        int `json:"id"`
	ProjectID int `json:"projectID"`
	UserID    int `json:"user_id"`
	//Duration    time.Duration `json:"duration"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	Description string    `json:"description"`
}
