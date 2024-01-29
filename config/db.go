package config

import (
	"graded-3/entity"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
) 

func InitDb() *gorm.DB {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := AutoMigrate(db); err != nil {
		panic(err)
	}
	return db
}

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&entity.User{}, &entity.UserActivityLog{}, &entity.Comment{}, &entity.Post{}); err != nil {
		panic(err)
	}
	return nil
}