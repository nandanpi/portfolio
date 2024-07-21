package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/public", "public")
	e.GET("/", s.HomePageHandler)

	e.GET("/admin", s.AdminPageHandler)
	e.GET("/admin/works", s.AdminWorksHandler)

	e.POST("/login", s.HandleLogin)

	e.GET("/getAllWorks", s.GetAllWorks)
	e.POST("/addWork", s.AddWorkHandler)
	e.POST("/toggleWorkPublished/:id", s.ToggleWorkPublished)
	return e
}
