package datastruct

import (
	"time"

	"github.com/google/uuid"
	pq "github.com/lib/pq"
)

type Practice struct {
	ID              uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey;" json:"id"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	Title           string         `gorm:"not null" json:"title"`
	Description     string         `gorm:"not null" json:"description"`
	Location        string         `gorm:"not null" json:"location"`
	Date            time.Time      `gorm:"not null" json:"date"`
	Expired         bool           `gorm:"default:false" json:"expired"`
	RequiredMembers int            `gorm:"not null" json:"required_members"`
	Members         pq.StringArray `gorm:"type:text[];not null" json:"members"`
	UserID          uuid.UUID      `gorm:"not null" json:"user_id"`
	User            User           `json:"user"`
}

func NewPractice(title, description, location string, userID uuid.UUID, requiredMembers int, date time.Time, members []string) *Practice {
	return &Practice{
		Title:           title,
		Description:     description,
		Location:        location,
		Date:            date,
		Expired:         false,
		UserID:          userID,
		RequiredMembers: requiredMembers,
		Members:         members,
	}
}
