package repository

import (
	"fmt"

	"github.com/NikhilSharma03/Yeager/internal/datastruct"
)

type UserRepository interface {
	Create(user *datastruct.User) error
}

type userRepository struct{}

func (*userRepository) Create(user *datastruct.User) error {
	result := DB.Create(user)
	if result.Error != nil {
		return fmt.Errorf("failed to create new user! %s", result.Error.Error())
	}
	return nil
}
