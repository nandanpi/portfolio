package server

import (
	"net/http"
	"portfolio/internal/auth"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/public", "public")
	e.GET("/", s.HomePageHndler)
	e.GET("/admin", auth.ProtectedRoute(s.AdminPageHandler))
	e.POST("/login", s.HandleLogin)

	return e
}
