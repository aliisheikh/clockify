package mysql

import (
	"clockify/users/domain"
	"clockify/users/infrastructure/entity"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type UserEpoImpl struct {
	DB *gorm.DB
}

func (u *UserEpoImpl) Delete(userID int) error {
	// Fetch the user by ID
	user := domain.User{ID: userID}
	result := u.DB.Where(&user).First(&user)
	if result.Error != nil {
		// Check if the user doesn't exist
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		// Return any other errors encountered during user retrieval
		return result.Error
	}

	// Delete the user
	result = u.DB.Delete(&user)
	if result.Error != nil {
		// Return any errors encountered during deletion
		return result.Error
	}

	// No error occurred, user deleted successfully
	return nil
}

func (u *UserEpoImpl) GetUserByID(userID int) (*domain.User, error) {
	if userID == 0 {
		// Return an error indicating that the user ID is invalid
		return nil, errors.New("invalid user ID")
	}

	var user domain.User
	result := u.DB.First(&user, userID)
	if result.Error != nil {
		// Check if the error is due to record not found
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		// Handle other errors
		return nil, result.Error
	}

	return &user, nil
}

func (u *UserEpoImpl) Save2(user *domain.User) (*domain.User, error) {
	result := u.DB.Save(user)
	fmt.Println("00000", user.ID)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (u *UserEpoImpl) Save(user *domain.User) error {
	// Check if the user already exists
	existingUser, err := u.GetUserByID(int(user.ID))
	if err != nil {
		// If the user does not exist, create a new record
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := u.DB.Create(user)
			if result.Error != nil {
				return result.Error
			}
			return nil
		}
		// Handle other errors
		return err
	}
	// Update existing user
	result := u.DB.Model(&existingUser).Updates(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *UserEpoImpl) GetAllUsers() ([]entity.User, error) {
	var users []entity.User

	if err := p.DB.Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// No users found
			return nil, fmt.Errorf("no users found")
		}
		// Other database error occurred
		return nil, fmt.Errorf("failed to retrieve users: %w", err)
	}

	return users, nil
}

func (u *UserEpoImpl) GetUserByUsername(username string) (*domain.User, error) {
	var user domain.User
	result := u.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

func NewUserEpoImpl(db *gorm.DB) domain.UserRepo {
	return &UserEpoImpl{DB: db}
}
