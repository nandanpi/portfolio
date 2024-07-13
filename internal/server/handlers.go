package server

import (
	"fmt"
	"net/http"
	"portfolio/internal/auth"
	"portfolio/views/home"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) HandleLogin(c echo.Context) error {

	username := c.FormValue("username")
	password := c.FormValue("password")

	dbUser, err := s.db.GetUser(username)
	fmt.Println(dbUser)
	if err != nil {
		return err
	}

	PassErr := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(password))
	if PassErr != nil {
		return PassErr
	}

	token, err := auth.GenerateJWT(dbUser.ID)
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.HttpOnly = true
	cookie.Expires = time.Now().Add(1 * time.Hour)
	// Only in Prod
	// cookie.Secure = true
	cookie.SameSite = http.SameSiteLaxMode
	c.SetCookie(cookie)

	return nil
}

func (s *Server) HomePageHndler(c echo.Context) error {
	return home.Index().Render(c.Request().Context(), c.Response())
}

func (s *Server) AdminPageHandler(c echo.Context) error {
	return home.Index().Render(c.Request().Context(), c.Response())
}
