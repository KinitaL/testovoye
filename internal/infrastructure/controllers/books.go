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
// @Summary Get all books
// @Description Retrieves a list of all books.
// @Tags books
// @Produce json
// @Success 200 {array} models.Book
// @Failure 500 {object} map[string]string "error"
// @Router /api/books [get]
func (c *Controller) GetAll(ctx echo.Context) error {
	books, err := c.u.GetAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, books)
}

// GetOne handles HTTP GET requests to retrieve a book by its ID.
// @Summary Get a single book
// @Description Retrieves a book by its unique ID.
// @Tags books
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} models.Book
// @Failure 400 {object} map[string]string "Invalid book ID"
// @Failure 404 {object} map[string]string "Book not found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /api/books/{id} [get]
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
// @Summary Create a new book
// @Description Adds a new book to the database.
// @Tags books
// @Accept json
// @Produce json
// @Param book body dto.CreateBookDto true "Book Data"
// @Success 200
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /api/books [post]
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
// @Summary Update an existing book
// @Description Modifies the details of an existing book.
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body models.Book true "Updated Book Data"
// @Success 200
// @Failure 400 {object} map[string]string "Invalid book ID / Invalid request body"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /api/books/{id} [patch]
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
// @Summary Delete a book
// @Description Removes a book from the database using its ID.
// @Tags books
// @Param id path string true "Book ID"
// @Success 200
// @Failure 400 {object} map[string]string "Invalid book ID"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /api/books/{id} [delete]
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
