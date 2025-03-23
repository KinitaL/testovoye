package controllers

import (
	"context"
	"github.com/KinitaL/testovoye/internal/infrastructure/controllers/dto"
	"github.com/KinitaL/testovoye/internal/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Controller struct handles HTTP requests and interacts with the usecase layer.
type (
	Controller struct {
		u usecase
	}

	// usecase defines the business logic layer interface for book operations.
	usecase interface {
		GetAll(ctx context.Context) ([]models.Book, error)                // Retrieves all books
		GetOne(ctx context.Context, ID uuid.UUID) (*models.Book, error)   // Retrieves a book by ID
		Create(ctx context.Context, book models.Book) error               // Creates a new book
		Update(ctx context.Context, ID uuid.UUID, book models.Book) error // Updates an existing book
		Delete(ctx context.Context, ID uuid.UUID) error                   // Deletes a book by ID
	}
)

// NewController initializes a new Controller instance.
func NewController(usecase usecase) *Controller {
	return &Controller{u: usecase}
}

// GetAll handles HTTP GET requests to retrieve all books.
func (c *Controller) GetAll(ctx echo.Context) error {
	books, err := c.u.GetAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, books)
}

// GetOne handles HTTP GET requests to retrieve a book by its ID.
func (c *Controller) GetOne(ctx echo.Context) error {
	ID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid book ID"})
	}
	book, err := c.u.GetOne(ctx.Request().Context(), ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if book == nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "book not found"})
	}
	return ctx.JSON(http.StatusOK, book)
}

// Create handles HTTP POST requests to create a new book.
func (c *Controller) Create(ctx echo.Context) error {
	var book dto.CreateBookDto
	if err := ctx.Bind(&book); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	if err := ctx.Validate(book); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	if err := c.u.Create(ctx.Request().Context(), models.Book{
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
	}); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.NoContent(http.StatusOK)
}

// Update handles HTTP PATCH requests to update an existing book.
func (c *Controller) Update(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid book ID"})
	}
	var book models.Book
	if err := ctx.Bind(&book); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	if err := ctx.Validate(book); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	if err := c.u.Update(ctx.Request().Context(), id, book); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.NoContent(http.StatusOK)
}

// Delete handles HTTP DELETE requests to remove a book by ID.
func (c *Controller) Delete(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid book ID"})
	}
	if err := c.u.Delete(ctx.Request().Context(), id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.NoContent(http.StatusOK)
}
