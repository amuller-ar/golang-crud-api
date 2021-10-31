package database

import (
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/domain"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Setup() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(
		&domain.Property{},
		&domain.User{},
	)

	if err != nil {
		log.Fatal(err)
	}

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
