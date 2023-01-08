package repository

import (
	"fmt"

	"github.com/NikhilSharma03/Yeager/internal/datastruct"
)

type PracticeRepository interface {
	Create(practice *datastruct.Practice) (*datastruct.Practice, error)
}

type practiceRepository struct{}

func (*practiceRepository) Create(practice *datastruct.Practice) (*datastruct.Practice, error) {
	result := DB.Create(practice)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create new practice! %s", result.Error.Error())
	}
	return practice, nil
}
