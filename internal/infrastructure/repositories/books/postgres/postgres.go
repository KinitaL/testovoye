package postgres

import (
	"context"
	"github.com/KinitaL/testovoye/internal/models"
	"github.com/KinitaL/testovoye/internal/usecases/books"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Repo is a GORM-based implementation of the book repository.
type Repo struct {
	db *gorm.DB
}

// NewPostgresRepo creates and returns a new repository instance using GORM and PostgreSQL.
func NewPostgresRepo(db *gorm.DB) books.Repository {
	return &Repo{db: db}
}

// GetAll retrieves all books from the database.
func (r *Repo) GetAll(ctx context.Context) ([]models.Book, error) {
	var rows []Book
	if err := r.db.WithContext(ctx).Find(&rows).Error; err != nil {
		return nil, err
	}
	result := make([]models.Book, len(rows))
	for i, book := range rows {
		result[i] = r.fromEntityToModel(book)
	}
	return result, nil
}

// GetOne retrieves a single book by its UUID.
func (r *Repo) GetOne(ctx context.Context, ID uuid.UUID) (*models.Book, error) {
	var book Book
	err := r.db.WithContext(ctx).First(&book, "id = ?", ID).Error
	if err != nil {
		return nil, err
	}
	model := r.fromEntityToModel(book)
	return &model, nil
}

// Create inserts a new book into the database.
func (r *Repo) Create(ctx context.Context, model models.Book) error {
	book := r.fromModelToEntity(model)
	return r.db.WithContext(ctx).Create(&book).Error
}

// Update modifies an existing book in the database.
func (r *Repo) Update(ctx context.Context, ID uuid.UUID, model models.Book) error {
	tx := r.db.WithContext(ctx).Begin()

	// Find existing book
	var existing Book
	if err := tx.First(&existing, "id = ?", ID).Error; err != nil {
		tx.Rollback()
		return err
	}

	book := r.fromModelToEntity(model)
	// Fill missing fields
	r.fillEmptyFields(&existing, &book)

	// Save updated book
	if err := tx.Save(&book).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Delete removes a book from the database by its UUID.
func (r *Repo) Delete(ctx context.Context, ID uuid.UUID) error {
	return r.db.WithContext(ctx).Where("id = ?", ID).Delete(&Book{}).Error
}

// fromEntityToModel converts an entity to a model (to the business logic layer from the db layer)
func (r *Repo) fromEntityToModel(entity Book) models.Book {
	return models.Book{
		ID:     entity.ID,
		Title:  entity.Title,
		Author: entity.Author,
		Year:   entity.Year,
	}
}

// fromModelToEntity converts a model to an entity (from the business logic layer to the db layer)
func (r *Repo) fromModelToEntity(model models.Book) Book {
	return Book{
		Base: Base{
			ID: model.ID,
		},
		Title:  model.Title,
		Author: model.Author,
		Year:   model.Year,
	}
}

// fillEmptyFields copies missing fields from the existing book to the updated book.
func (r *Repo) fillEmptyFields(existing, updated *Book) {
	if updated.Title == "" {
		updated.Title = existing.Title
	}
	if updated.Author == "" {
		updated.Author = existing.Author
	}
	if updated.Year == 0 {
		updated.Year = existing.Year
	}
}
