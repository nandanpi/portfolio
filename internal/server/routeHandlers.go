package server

import (
	"net/http"
	"portfolio/views/admin"
	"portfolio/views/home"

	"github.com/labstack/echo/v4"
)

func (s *Server) HomePageHandler(c echo.Context) error {
	return home.Index().Render(c.Request().Context(), c.Response())
}

func (s *Server) AdminPageHandler(c echo.Context) error {
	return admin.Index().Render(c.Request().Context(), c.Response())
}

func (s *Server) AdminWorksHandler(c echo.Context) error {

	works, err := s.db.GetAllWorks()
	if err != nil {
		c.Logger().Error("Failed to fetch works: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch works",
		})
	}

	return admin.Works(works).Render(c.Request().Context(), c.Response())
}
