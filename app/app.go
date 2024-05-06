package app

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	e *echo.Echo
}

func (s *Server) Init() error {
	s.e = echo.New()

	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())
	s.e.Use(s.HeaderMiddleware())

	s.e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	s.e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	return nil
}

func (s *Server) Start() error {
	return s.e.Start(":8080")
}

func (s *Server) Stop() error {
	return s.e.Close()
}

func (s *Server) HeaderMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("dc", os.Getenv("NAME"))
			return next(c)
		}
	}
}
