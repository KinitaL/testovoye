package server

import (
	"github.com/KinitaL/testovoye/config"
	"github.com/KinitaL/testovoye/pkg/validator"
	"github.com/labstack/echo/v4"
	"net"
)

// BuildServer returns instance of http server we are using in all controllers.
func BuildServer(cfg config.Service, middlewares ...echo.MiddlewareFunc) *echo.Echo {
	e := echo.New()

	l, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		e.Logger.Fatal(err)
	}
	e.Listener = l
	e.Validator = validator.New()

	if len(middlewares) > 0 {
		e.Use(middlewares...)
	}

	return e
}
