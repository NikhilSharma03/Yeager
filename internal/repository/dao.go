package repository

import (
	"fmt"

	"github.com/NikhilSharma03/Yeager/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// The DB variable will contain the database client
// It is initialized when a new DAO is create
var DB *gorm.DB

// DAO interface
type DAO interface {
	NewPracticeRepository() PracticeRepository
}

// The dao struct implement the DAO interface
type dao struct{}

// The NewDAO function returns a pointer to dao struct which implement DAO interface
func NewDAO(db *gorm.DB) DAO {
	DB = db
	return &dao{}
}

func (*dao) NewPracticeRepository() PracticeRepository {
	return &practiceRepository{}
}

// ConnectDB connects to postgres Database using provided config
func ConnectDB(cfg *config.Config) (*gorm.DB, error) {
	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.DBHost, cfg.DBUserName, cfg.DBUserPassword, cfg.DBName, cfg.DBPort)
	return gorm.Open(postgres.Open(dbURI), &gorm.Config{})
}
