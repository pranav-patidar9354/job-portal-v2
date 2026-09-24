package config

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/pranav-patidar9354/job-portal-v2/backend/internal/models"
)

var DB *gorm.DB

func ConnectDatabase() error {
	host := env("DB_HOST", "localhost")
	port := env("DB_PORT", "3306")
	user := env("DB_USER", "root")
	password := os.Getenv("DB_PASSWORD")
	name := env("DB_NAME", "job_portal_v2")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := db.AutoMigrate(
		&models.User{},
		&models.Job{},
		&models.Application{},
		&models.SavedJob{},
	); err != nil {
		return err
	}

	DB = db
	return nil
}

func env(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
