package config

import (
	"os"

	"github.com/gabriel-O-C/godin/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeDb() (*gorm.DB, error) {
	logger := GetLogger("database")
	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("database file not found. Creating one...")

		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			return nil, err
		}

		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("database init error %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("database migration error %v", err)
		return nil, err
	}

	return db, nil
}
