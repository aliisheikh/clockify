package entity

import (
	"clockify/project/domain"
	domain2 "clockify/users/domain"
	"time"
)

type TimeEntry struct {
	ID          int        `gorm:"primary_key;autoIncrement" json:"id"`
	ProjectID   *int       `gorm:"not null;column:project_id" json:"projectID"`
	UserID      int        `gorm:"not null;column:user_id" json:"user_id;foreignkey:UserID;constraint:OnDelete:CASCADE;" binding:"max=1000;column:user_id"`
	StartTime   time.Time  `gorm:"type:timestamp;not null;column:StartTime" json:"startTime"`
	EndTime     *time.Time `gorm:"type:timestamp;not null;column:EndTime" json:"endTime"`
	Description string     `gorm:"type:text;column:Description" json:"description,omitempty"`

	Project domain.Projects `gorm:"foreignKey:ProjectID;references:ID"`
	User    domain2.User    `gorm:"foreignKey:UserID;references:ID"`
}

func (TimeEntry) TableName() string {
	return "time_entries"
}
