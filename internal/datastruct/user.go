package datastruct

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"not null;unique_index" json:"email"`
	Password  string    `json:"password"`
	Avatar    string    `gorm:"not null" json:"avatar"`
}

func NewUser(name, email, password, avatar string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
		Avatar:   avatar,
	}
}
