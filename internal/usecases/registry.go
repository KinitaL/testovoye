package usecases

import "github.com/KinitaL/testovoye/internal/usecases/books"

type (
	Registry struct {
		Books books.Books
	}
	RepositoriesRegistry struct {
		Books books.Repository
	}
)

func NewRegistry(repos *RepositoriesRegistry) *Registry {
	return &Registry{
		Books: books.NewBooksUsecase(repos.Books),
	}
}

func NewRepositoriesRegistry(books books.Repository) *RepositoriesRegistry {
	return &RepositoriesRegistry{Books: books}
}
