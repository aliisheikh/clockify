package domain

import (
	"clockify/project/domain"
	domain2 "clockify/timeEntry/domain"
	"clockify/users/infrastructure/entity"
)

type User struct {
	ID        int                 `json:"id"`
	Username  string              `json:"username"`
	Email     string              `json:"email"`
	Password  string              `json:"password"`
	projects  []domain.Projects   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	timeEntry []domain2.TimeEntry `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// Add other fields as needed
}

type UserRepo interface {
	Save(user *User) error
	Delete(user int) error
	Save2(user *User) (*User, error)
	GetUserByID(userId int) (*User, error)
	GetUserByUsername(username string) (*User, error)
	GetAllUsers() ([]entity.User, error)
}
