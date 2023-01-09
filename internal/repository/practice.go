package repository

import (
	"errors"
	"fmt"

	"github.com/NikhilSharma03/Yeager/internal/datastruct"
	"gorm.io/gorm"
)

type PracticeRepository interface {
	Create(practice *datastruct.Practice) error
	GetAll(practices *[]datastruct.Practice) error
	GetByUserID(practices *[]datastruct.Practice, userID string) error
	GetByID(practice *datastruct.Practice, id string) error
	UpdateByID(practice *datastruct.Practice, id string) error
	DeleteByID(practice *datastruct.Practice, id string) error
}

type practiceRepository struct{}

var (
	errNoPracticeRecordsFound = errors.New("found no records for practices")
	errNoPracticeFoundByID    = errors.New("found no practice by provided id")
)

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
	if result.RowsAffected == 0 {
		return errNoPracticeRecordsFound
	}
	return nil
}

func (*practiceRepository) GetByUserID(practices *[]datastruct.Practice, userID string) error {
	result := DB.Find(practices, "creator=?", userID)
	if result.Error != nil {
		return fmt.Errorf("failed to fetch all practice by user id! %s", result.Error.Error())
	}
	return nil
}

func (*practiceRepository) GetByID(practice *datastruct.Practice, id string) error {
	result := DB.First(practice, "id=?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errNoPracticeFoundByID
		}
		return fmt.Errorf("failed to fetch practice by id! %s", result.Error.Error())
	}
	return nil
}

func (pr *practiceRepository) UpdateByID(practice *datastruct.Practice, id string) error {
	result := DB.Model(practice).Where("id=?", id).Updates(practice)
	if result.Error != nil {
		return fmt.Errorf("failed to update practice by id! %s", result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return errNoPracticeFoundByID
	}
	return nil
}

func (pr *practiceRepository) DeleteByID(practice *datastruct.Practice, id string) error {
	result := DB.Where("id=?", id).Delete(practice)
	if result.Error != nil {
		return fmt.Errorf("failed to delete practice by id! %s", result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return errNoPracticeFoundByID
	}
	return nil
}
