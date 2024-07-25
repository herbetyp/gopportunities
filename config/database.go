package config

import (
	"github.com/herbetyp/gopportunities/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	logger := GetLogger("database")

	dsn := "host=localhost user=postgres password=s3cr3t dbname=openings_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.ErrorF("Opening connection error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&models.Opening{})
	if err != nil {
		logger.ErrorF("Database migration error: %v", err)
		return nil, err
	}

	return db, nil
}
