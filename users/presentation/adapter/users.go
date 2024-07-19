package adapter

import (
	"clockify/users/domain"
	"clockify/users/presentation/models"
)

func DomainToUser(user domain.User) models.User {
	return models.User{
		ID:       int64(user.ID),
		Username: user.Username,
		Email:    user.Email,
		//Password: user.Password,
	}
}
