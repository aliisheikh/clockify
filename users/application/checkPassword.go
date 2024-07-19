package application

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func (u *UserServiceImp) CheckPassword(username, password string) (bool, error) {

	if username == "" {
		return false, errors.New("username can't be empty")
	}
	if password == "" {
		return false, errors.New("password can't be empty")
	}

	user, err := u.usersRepository.GetUserByUsername(username)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, nil
	}

	return true, nil
}
