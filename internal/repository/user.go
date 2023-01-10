package repository

import (
	"errors"
	"fmt"

	"github.com/NikhilSharma03/Yeager/internal/datastruct"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *datastruct.User) error
	GetByID(user *datastruct.User, id string) error
	UpdateByID(user *datastruct.User, id string) error
	DeleteByID(user *datastruct.User, id string) error
}

type userRepository struct{}

var (
	errNoUserFoundByID = errors.New("found no user by provided id")
)

func (*userRepository) Create(user *datastruct.User) error {
	result := DB.Create(user)
	if result.Error != nil {
		return fmt.Errorf("failed to create new user! %s", result.Error.Error())
	}
	return nil
}

func (*userRepository) GetByID(user *datastruct.User, id string) error {
	result := DB.First(user, "id=?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errNoUserFoundByID
		}
		return fmt.Errorf("failed to fetch user by id! %s", result.Error.Error())
	}
	return nil
}

func (*userRepository) UpdateByID(user *datastruct.User, id string) error {
	result := DB.Model(user).Where("id=?", id).Updates(user)
	if result.Error != nil {
		return fmt.Errorf("failed to update user by id! %s", result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return errNoUserFoundByID
	}
	return nil
}

func (*userRepository) DeleteByID(user *datastruct.User, id string) error {
	result := DB.Where("id=?", id).Delete(user)
	if result.Error != nil {
		return fmt.Errorf("failed to delete user by id! %s", result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return errNoUserFoundByID
	}
	return nil
}
