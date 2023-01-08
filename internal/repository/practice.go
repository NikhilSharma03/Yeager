package repository

import (
	"fmt"

	"github.com/NikhilSharma03/Yeager/internal/datastruct"
)

type PracticeRepository interface {
	Create(practice *datastruct.Practice) error
	GetAll(practices *[]datastruct.Practice) error
}

type practiceRepository struct{}

func (*practiceRepository) Create(practice *datastruct.Practice) error {
	result := DB.Create(practice)
	if result.Error != nil {
		return fmt.Errorf("failed to create new practice! %s", result.Error.Error())
	}
	return nil
}

func (*practiceRepository) GetAll(practices *[]datastruct.Practice) error {
	result := DB.Find(practices)
	if result.Error != nil {
		return fmt.Errorf("failed to fetch all practice! %s", result.Error.Error())
	}
	return nil
}
