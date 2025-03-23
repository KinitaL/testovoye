package books

import (
	"context"
	"github.com/KinitaL/testovoye/internal/models"
	"github.com/google/uuid"
)

//go:generate go install go.uber.org/mock/mockgen@v0.5.0
//go:generate mockgen -destination usecase_mock.go -package books . Books

// Books interface defines the main operations for managing books.
type (
	Books interface {
		GetAll(ctx context.Context) ([]models.Book, error)                // Retrieve all books
		GetOne(ctx context.Context, ID uuid.UUID) (*models.Book, error)   // Get a single book by ID
		Create(ctx context.Context, book models.Book) error               // Create a new book
		Update(ctx context.Context, ID uuid.UUID, book models.Book) error // Update an existing book
		Delete(ctx context.Context, ID uuid.UUID) error                   // Delete a book by ID
	}

	// books struct implements the Books interface.
	books struct {
		repo Repository // Repository for data operations
	}
)

// NewBooksUsecase creates and returns a new instance of the book use case.
func NewBooksUsecase(repo Repository) Books {
	return &books{
		repo: repo,
	}
}

// GetAll retrieves a list of all books.
func (u *books) GetAll(ctx context.Context) ([]models.Book, error) {
	return u.repo.GetAll(ctx)
}

// GetOne fetches a book by its ID.
func (u *books) GetOne(ctx context.Context, ID uuid.UUID) (*models.Book, error) {
	return u.repo.GetOne(ctx, ID)
}

// Create adds a new book with a unique identifier.
func (u *books) Create(ctx context.Context, book models.Book) error {
	book.ID = uuid.New() // Generate a new UUID for the book
	return u.repo.Create(ctx, book)
}

// Update modifies an existing book by its ID.
func (u *books) Update(ctx context.Context, ID uuid.UUID, book models.Book) error {
	book.ID = ID // Ensure the ID remains unchanged
	return u.repo.Update(ctx, ID, book)
}

// Delete removes a book by its ID.
func (u *books) Delete(ctx context.Context, ID uuid.UUID) error {
	return u.repo.Delete(ctx, ID)
}
