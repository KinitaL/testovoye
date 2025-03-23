package books

import (
	"context"
	"github.com/KinitaL/testovoye/internal/models"
	"github.com/google/uuid"
)

//go:generate mockgen -destination repository_mock.go -package books . Repository

type Repository interface {
	GetAll(ctx context.Context) ([]models.Book, error)
	GetOne(ctx context.Context, ID uuid.UUID) (*models.Book, error)
	Create(ctx context.Context, book models.Book) error
	Update(ctx context.Context, ID uuid.UUID, book models.Book) error
	Delete(ctx context.Context, ID uuid.UUID) error
}
