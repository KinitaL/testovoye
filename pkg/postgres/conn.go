package postgres

import (
	"fmt"
	"github.com/KinitaL/testovoye/config"
	repo "github.com/KinitaL/testovoye/internal/infrastructure/repositories/books/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewPostgresDB initializes a PostgreSQL database connection using GORM.
func NewPostgresDB(cfg config.DB) (*gorm.DB, error) {
	// Format DSN (Data Source Name)
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Enable query logging
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Run auto-migration
	if err := db.AutoMigrate(&repo.Book{}); err != nil {
		return nil, fmt.Errorf("failed to migrate: %w", err)
	}

	return db, nil
}
