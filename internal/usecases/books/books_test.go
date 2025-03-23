package books

import (
	"context"
	"github.com/KinitaL/testovoye/internal/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestCreate(t *testing.T) {
	// init mocks
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repo := NewMockRepository(mockCtrl)

	// init core
	usecase := NewBooksUsecase(repo)

	// test cases
	cases := []struct {
		name string

		req models.Book
		err error
	}{
		{
			name: "Create",
			req: models.Book{
				Title:  "Test Book",
				Author: "Tester",
				Year:   2025,
			},
			err: nil,
		},
	}

	// execution
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			ctx := context.Background()
			repo.EXPECT().Create(ctx, gomock.Any()).Return(testCase.err).AnyTimes()
			// execution
			err := usecase.Create(ctx, testCase.req)
			assert.Equal(t, testCase.err, err)
		})
	}
}

func TestGetAll(t *testing.T) {
	// init mocks
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repo := NewMockRepository(mockCtrl)

	// init core
	usecase := NewBooksUsecase(repo)

	// test cases
	cases := []struct {
		name string

		resp []models.Book
		err  error
	}{
		{
			name: "GetAll",

			resp: []models.Book{
				{
					ID:     uuid.New(),
					Title:  "Test Book",
					Author: "Test Author",
					Year:   2024,
				},
			},
			err: nil,
		},
	}

	// execution
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			ctx := context.Background()
			repo.EXPECT().GetAll(ctx).Return(testCase.resp, testCase.err).AnyTimes()
			// execution
			resp, err := usecase.GetAll(ctx)
			assert.Equal(t, testCase.err, err)
			assert.Equal(t, testCase.resp, resp)
		})
	}
}

func TestGetOne(t *testing.T) {
	// init mocks
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repo := NewMockRepository(mockCtrl)

	// init core
	usecase := NewBooksUsecase(repo)

	// test cases
	cases := []struct {
		name string

		req  uuid.UUID
		resp *models.Book
		err  error
	}{
		{
			name: "GetOne",

			req: uuid.New(),
			resp: &models.Book{
				ID:     uuid.New(),
				Title:  "New Test",
				Author: "Tester",
				Year:   2024,
			},
			err: nil,
		},
	}

	// execution
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			ctx := context.Background()
			repo.EXPECT().GetOne(ctx, testCase.req).Return(testCase.resp, testCase.err).AnyTimes()
			// execution
			resp, err := usecase.GetOne(ctx, testCase.req)
			assert.Equal(t, testCase.err, err)
			assert.Equal(t, testCase.resp, resp)
		})
	}
}

func TestUpdate(t *testing.T) {
	// init mocks
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repo := NewMockRepository(mockCtrl)

	// init core
	usecase := NewBooksUsecase(repo)

	// test cases
	cases := []struct {
		name string

		req models.Book
		ID  uuid.UUID
		err error
	}{
		{
			name: "Update full",

			ID: uuid.New(),
			req: models.Book{
				Title:  "Test Update",
				Author: "Tester",
				Year:   2025,
			},
			err: nil,
		},
		{
			name: "Update one field",

			ID: uuid.New(),
			req: models.Book{
				Title: "Test Update",
			},
			err: nil,
		},
	}

	// execution
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			ctx := context.Background()
			repo.EXPECT().Update(ctx, testCase.ID, gomock.Any()).Return(testCase.err).AnyTimes()
			// execution
			err := usecase.Update(ctx, testCase.ID, testCase.req)
			assert.Equal(t, testCase.err, err)
		})
	}
}

func TestDelete(t *testing.T) {
	// init mocks
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repo := NewMockRepository(mockCtrl)

	// init core
	usecase := NewBooksUsecase(repo)

	// test cases
	cases := []struct {
		name string

		req uuid.UUID
		err error
	}{
		{
			name: "GetOne",

			req: uuid.New(),
			err: nil,
		},
	}

	// execution
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			ctx := context.Background()
			repo.EXPECT().Delete(ctx, testCase.req).Return(testCase.err).AnyTimes()
			// execution
			err := usecase.Delete(ctx, testCase.req)
			assert.Equal(t, testCase.err, err)
		})

	}
}
