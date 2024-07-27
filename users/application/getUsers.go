package application

import (
	"clockify/users/domain"
	"fmt"
)

func (u *UserServiceImp) GetUserByID(userID int) (*domain.User, error) {
	userData, err := u.usersRepository.GetUserByID(userID)
	fmt.Println(userData)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %v", err)
	}

	if userData == nil {
		return nil, fmt.Errorf("user with ID %d not found", userID)
	}

	userResponse := &domain.User{
		ID:       userData.ID,
		Email:    userData.Email,
		Username: userData.Username,
		// Avoid returning password in response
	}

	return userResponse, nil
}
