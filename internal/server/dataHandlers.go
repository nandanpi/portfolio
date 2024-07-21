package server

import (
	"fmt"
	"net/http"
	"portfolio/internal/auth"
	"portfolio/internal/types"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) AddWorkHandler(c echo.Context) error {

	work := types.AddWorkReq{}
	work.Title = c.FormValue("title")
	work.Description = c.FormValue("description")
	work.Image = c.FormValue("image")
	work.GithubLink = c.FormValue("github_link")
	work.DemoLink = c.FormValue("demo_link")
	techStackStr := c.FormValue("tech_stack")
	techStack := []uint{}
	for _, str := range strings.Split(techStackStr, ",") {
		val, err := strconv.ParseUint(str, 10, 32)
		if err != nil {
			return err
		}
		techStack = append(techStack, uint(val))
	}
	work.TechStack = techStack

	err := s.db.AddWork(work)
	if err != nil {
		return err
	}

	return nil

}

func (s *Server) ToggleWorkPublished(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return err
	}

	err = s.db.ToggleWorkPublished(uint(id))
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) DeleteWork(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return err
	}

	err = s.db.DeleteWork(uint(id))
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) GetAllWorks(c echo.Context) error {

	works, err := s.db.GetAllWorks()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, works)
}

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
