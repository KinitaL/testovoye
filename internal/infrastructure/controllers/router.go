package controllers

import (
	_ "github.com/KinitaL/testovoye/docs"
	"github.com/KinitaL/testovoye/internal/usecases"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

//go:generate go install github.com/swaggo/swag/cmd/swag@latest
//go:generate swag init -o ../../../docs/swagger/api -g router.go --parseInternal --parseDependency --instanceName api

// Register
// @title Books API
// @version 1.0
// @description Сервис книг
// @basePath /api
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

	server.GET("/swagger/*", echoSwagger.WrapHandler)
}
