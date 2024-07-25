package config

import (
	"fmt"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func InitConfig() error {
	var err error

	db, err = InitDatabase()
	if err != nil {
		return fmt.Errorf("Error initializing database: %v", err)
	}
	return nil
}

func GetDatabase() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}
