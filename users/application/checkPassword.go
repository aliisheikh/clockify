package application

import (
	"clockify/users/domain"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func (u *UserServiceImp) CheckPassword(username, password string) (*domain.User, error) {

	if username == "" {
		return nil, errors.New("username can't be empty")
	}
	if password == "" {
		return nil, errors.New("password can't be empty")
	}

	user, err := u.usersRepository.GetUserByUsername(username)
	fmt.Println(user)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}
