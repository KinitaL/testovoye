package dto

type (
	CreateBookDto struct {
		Title  string `json:"title" validate:"required"`
		Author string `json:"author" validate:"required"`
		Year   uint16 `json:"year" validate:"required"`
	}
	UpdateBookDto struct {
		Title  string `json:"title,omitempty"`
		Author string `json:"author,omitempty"`
		Year   uint16 `json:"year,omitempty"`
	}
)
