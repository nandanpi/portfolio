package server

import (
	"encoding/json"
	"net/http"
	"portfolio/internal/auth"
	"portfolio/internal/types"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) HandleLogin(c echo.Context) error {
	loginReq := &types.AuthRequest{}
	json.NewDecoder(c.Request().Body).Decode(loginReq)

	dbUser := s.db.GetUser(c, loginReq.Username)

	resp := make(map[string]string)

	PassErr := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginReq.Password))
	if PassErr != nil {
		resp["message"] = "Wrong Password"
		return c.JSON(http.StatusUnauthorized, resp)

	}

	token, err := auth.GenerateJWT(dbUser.ID)
	if err != nil {
		resp["message"] = "Error generating JWT Token"
		return c.JSON(http.StatusUnauthorized, resp)

	}

	return c.JSON(http.StatusOK, map[string]interface{}{"token": token, "user": map[string]interface{}{"username": dbUser.Username, "id": dbUser.ID}})
}
