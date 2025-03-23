package models

import "github.com/google/uuid"

// Book is a model that is used as a business logic unit
type Book struct {
	ID     uuid.UUID
	Title  string
	Author string
	Year   uint16
}
