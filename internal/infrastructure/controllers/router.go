package controllers

import (
	"github.com/KinitaL/testovoye/internal/usecases"
	"github.com/labstack/echo/v4"
)

func Register(server *echo.Echo, registry *usecases.Registry) {

	api := server.Group("/api")

	{
		books := NewController(registry.Books)
		api.POST("/books", books.Create)
		api.GET("/books", books.GetAll)
		api.GET("/books/:id", books.GetOne)
		api.PATCH("/books/:id", books.Update)
		api.DELETE("/books/:id", books.Delete)
	}
}
