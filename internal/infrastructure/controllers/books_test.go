package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/KinitaL/testovoye/internal/infrastructure/controllers/dto"
	"github.com/KinitaL/testovoye/internal/models"
	usecase_mock "github.com/KinitaL/testovoye/internal/usecases/books"
	"github.com/KinitaL/testovoye/pkg/validator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGetAll tests the GetAll controller method
func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	mockUsecase := usecase_mock.NewMockBooks(ctrl)
	controller := NewController(mockUsecase)

	books := []models.Book{
		{ID: uuid.New(), Title: "Book 1", Author: "Author 1", Year: 2021},
		{ID: uuid.New(), Title: "Book 2", Author: "Author 2", Year: 2022},
	}

	mockUsecase.EXPECT().GetAll(gomock.Any()).Return(books, nil).AnyTimes()

	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	err := controller.GetAll(ctx)
	assert.Equal(t, err, nil)
	assert.Equal(t, rec.Code, http.StatusOK)

	var response []models.Book
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(response), len(books))
}

// TestGetOne tests GetOne with a valid and invalid UUID
func TestGetOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	mockUsecase := usecase_mock.NewMockBooks(ctrl)
	controller := NewController(mockUsecase)

	bookID := uuid.New()
	book := &models.Book{ID: bookID, Title: "Test Book", Author: "Test Author", Year: 2023}

	t.Run("Success", func(t *testing.T) {
		mockUsecase.EXPECT().GetOne(gomock.Any(), bookID).Return(book, nil).AnyTimes()

		req := httptest.NewRequest(http.MethodGet, "/books/"+bookID.String(), nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("id")
		ctx.SetParamValues(bookID.String())

		err := controller.GetOne(ctx)
		assert.Equal(t, err, nil)
		assert.Equal(t, rec.Code, http.StatusOK)

		var response models.Book
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, err, nil)
		assert.Equal(t, response.ID, book.ID)
	})

	t.Run("Invalid UUID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/books/invalid-uuid", nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("id")
		ctx.SetParamValues("invalid-uuid")

		err := controller.GetOne(ctx)
		assert.Equal(t, err, nil)
		assert.Equal(t, rec.Code, http.StatusBadRequest)
	})

	t.Run("Not Found", func(t *testing.T) {
		bookID := uuid.New()
		mockUsecase.EXPECT().GetOne(gomock.Any(), bookID).Return(nil, nil).AnyTimes()

		req := httptest.NewRequest(http.MethodGet, "/books/"+bookID.String(), nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("id")
		ctx.SetParamValues(bookID.String())

		err := controller.GetOne(ctx)
		assert.Equal(t, err, nil)
		assert.Equal(t, rec.Code, http.StatusNotFound)
	})
}

// TestCreate tests the Create controller method
func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	e.Validator = validator.New()
	mockUsecase := usecase_mock.NewMockBooks(ctrl)
	controller := NewController(mockUsecase)

	payload := dto.CreateBookDto{Title: "New Book", Author: "New Author", Year: 2023}
	book := models.Book{Title: payload.Title, Author: payload.Author, Year: payload.Year}

	mockUsecase.EXPECT().Create(gomock.Any(), book).Return(nil).AnyTimes()

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	err := controller.Create(ctx)
	assert.Equal(t, err, nil)
	assert.Equal(t, rec.Code, http.StatusOK)
}

// TestUpdate tests the Update controller method
func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	e.Validator = validator.New()
	mockUsecase := usecase_mock.NewMockBooks(ctrl)
	controller := NewController(mockUsecase)

	bookID := uuid.New()
	book := models.Book{Title: "Updated Title", Author: "Updated Author", Year: 2024}

	mockUsecase.EXPECT().Update(gomock.Any(), bookID, book).Return(nil).AnyTimes()

	body, _ := json.Marshal(book)
	req := httptest.NewRequest(http.MethodPatch, "/books/"+bookID.String(), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues(bookID.String())

	err := controller.Update(ctx)
	assert.Equal(t, err, nil)
	assert.Equal(t, rec.Code, http.StatusOK)
}

// TestDelete tests the Delete controller method
func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	mockUsecase := usecase_mock.NewMockBooks(ctrl)
	controller := NewController(mockUsecase)

	bookID := uuid.New()

	mockUsecase.EXPECT().Delete(gomock.Any(), bookID).Return(nil).AnyTimes()

	req := httptest.NewRequest(http.MethodDelete, "/books/"+bookID.String(), nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues(bookID.String())

	err := controller.Delete(ctx)
	assert.Equal(t, err, nil)
	assert.Equal(t, rec.Code, http.StatusOK)
}
