package domain

import "clockify/users/infrastructure/entity"

type UserService interface {
	Create(user User) (*User, error)
	Delete(userId int) error
	GetUserByID(userId int) (*User, error)
	CheckPassword(username, password string) (*User, error)
	GetAllUsers() ([]entity.User, error)
	GenerateToken(userID int, username string) (string, error)
}
