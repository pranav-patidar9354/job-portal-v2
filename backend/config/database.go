package config

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/pranav-patidar9354/job-portal-v2/backend/internal/models"
)

var DB *gorm.DB

func ConnectDatabase() error {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		host := env("DB_HOST", "localhost")
		port := env("DB_PORT", "5432")
		user := env("DB_USER", "postgres")
		password := os.Getenv("DB_PASSWORD")
		name := env("DB_NAME", "job_portal_v2")
		sslmode := env("DB_SSLMODE", "disable")

		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			host, user, password, name, port, sslmode,
		)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
