package database

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (d *service) GetUser(c echo.Context, username string) *User {
	user := &User{}
	resp := make(map[string]string)
	result := d.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		resp["message"] = "User not found"
		c.JSON(http.StatusNotFound, resp)
		return nil
	}
	return user
}
