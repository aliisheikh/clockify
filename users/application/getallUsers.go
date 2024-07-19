package application

import (
	"clockify/users/infrastructure/entity"
)

func (s *UserServiceImp) GetAllUsers() ([]entity.User, error) {

	return s.usersRepository.GetAllUsers()
}
