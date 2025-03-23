package postgres

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type (
	// Base contains common columns for all tables.
	Base struct {
		ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt `sql:"index"`
	}

	// Book contains columns for books table
	Book struct {
		Base
		Title  string `gorm:"not_null"`
		Author string `gorm:"not_null"`
		Year   uint16 `gorm:"type:int;not_null"`
	}
)
