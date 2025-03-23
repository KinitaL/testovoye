package books

import (
	"context"
	"fmt"
	"github.com/KinitaL/testovoye/internal/models"
	"github.com/KinitaL/testovoye/internal/usecases/books"
	"github.com/google/uuid"
	"sync"
)

// InMemoryRepo is a thread-safe in-memory implementation of the book repository.
type InMemoryRepo struct {
	sync.RWMutex
	books map[uuid.UUID]models.Book // Map to store books using UUID as the key
}

// NewInMemoryRepo creates and returns a new instance of InMemoryRepo.
func NewInMemoryRepo() books.Repository {
	return &InMemoryRepo{
		RWMutex: sync.RWMutex{},
		books:   make(map[uuid.UUID]models.Book),
	}
}

// GetAll retrieves all books from the repository.
func (r *InMemoryRepo) GetAll(_ context.Context) ([]models.Book, error) {
	result := make([]models.Book, 0, len(r.books))

	r.RLock()
	defer r.RUnlock()

	for _, b := range r.books {
		result = append(result, b)
	}
	return result, nil
}

// GetOne retrieves a single book by its UUID.
func (r *InMemoryRepo) GetOne(_ context.Context, ID uuid.UUID) (*models.Book, error) {
	r.RLock()
	defer r.RUnlock()

	book, ok := r.books[ID]
	if !ok {
		return nil, fmt.Errorf("book with ID = %s doesn't exist", ID)
	}
	return &book, nil
}

// Create adds a new book to the repository.
func (r *InMemoryRepo) Create(_ context.Context, book models.Book) error {
	r.Lock()
	defer r.Unlock()
	r.books[book.ID] = book
	return nil
}

// Update modifies an existing book in the repository.
func (r *InMemoryRepo) Update(_ context.Context, ID uuid.UUID, book models.Book) error {
	r.Lock()
	defer r.Unlock()

	old, ok := r.books[ID]
	if !ok {
		return fmt.Errorf("book with ID = %s doesn't exist", ID)
	}

	r.fillEmptyFields(&old, &book)
	r.books[ID] = book
	return nil
}

// Delete removes a book from the repository by its UUID.
func (r *InMemoryRepo) Delete(_ context.Context, ID uuid.UUID) error {
	r.Lock()
	defer r.Unlock()

	delete(r.books, ID)
	return nil
}

// fillEmptyFields copies missing fields from the old book to the new one.
func (r *InMemoryRepo) fillEmptyFields(old, new *models.Book) {
	if new.Title == "" {
		new.Title = old.Title
	}
	if new.Author == "" {
		new.Author = old.Author
	}
	if new.Year == 0 {
		new.Year = old.Year
	}
}
