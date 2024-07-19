package application

import (
	"errors"
	"gorm.io/gorm"
)

//
//type UserServiceImp struct {
//	usersRepository mysql.UserEpoImpl
//}
//
//func NewUserServiceImp(usersRepository mysql.UserEpoImpl) *UserServiceImp {
//	return &UserServiceImp{
//		usersRepository: usersRepository,
//	}
//}

func (u *UserServiceImp) Delete(userID int) error {
	// Check if the user with the given ID exists

	_, err := u.usersRepository.GetUserByID(userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// If the user doesn't exist, return nil indicating successful deletion
		return nil
	} else if err != nil {
		// If an error occurred while retrieving the user, return the error
		return err
	}

	// Attempt to delete the user
	err = u.usersRepository.Delete(userID)
	if err != nil {
		// If an error occurred while deleting the user, return the error
		return err
	}

	// User successfully deleted
	return nil
}
