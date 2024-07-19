package application

import (
	"clockify/users/domain"
	"golang.org/x/crypto/bcrypt"

	"errors"
	"fmt"
)

type UserServiceImp struct {
	usersRepository domain.UserRepo
	jwtKey          []byte
}

func NewUserServiceImp(usersRepository domain.UserRepo, jwtKey []byte) *UserServiceImp {
	return &UserServiceImp{
		usersRepository: usersRepository,
		jwtKey:          jwtKey,
	}
}

func (u *UserServiceImp) Create(user domain.User) (*domain.User, error) {
	// Validate required fields
	if user.Username == "" {
		return &user, errors.New("username is required")
	}
	if user.Email == "" {
		return &user, errors.New("email is required")
	}
	if user.Password == "" {
		return &user, errors.New("password is required")
	}

	// Hash the password
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return &user, fmt.Errorf("failed to hash password: %w", err)
	}

	// Check if email already exists
	existingUsers, err := u.usersRepository.GetAllUsers()
	if err != nil {
		return &user, fmt.Errorf("failed to retrieve users: %w", err)
	}
	for _, existingUser := range existingUsers {
		if existingUser.Email == user.Email {
			return &user, errors.New("email already exists")
		}
	}

	userModel := &domain.User{
		Email:    user.Email,
		Username: user.Username,
		Password: hashedPassword,
	}
	savedUser, err := u.usersRepository.Save2(userModel)
	if err != nil {
		return &user, fmt.Errorf("failed to save user: %w", err)
	}

	user.ID = savedUser.ID
	// Return the saved user data
	return savedUser, nil
}

// hashPassword hashes the password using bcrypt with a cost factor of 12
func hashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
