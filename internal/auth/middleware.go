package auth

import (
	"os"
	"portfolio/views/login"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

var JWTKey = []byte(os.Getenv("JWT_KEY"))

func GenerateJWT(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTKey)
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte(JWTKey)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}

func ProtectedRoute(f func(c echo.Context) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("token")
		if err != nil {
			// return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
			return login.Login().Render(c.Request().Context(), c.Response())
		}
		tokenString := cookie.Value
		_, err = verifyToken(tokenString)
		if err != nil {
			// return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
			return login.Login().Render(c.Request().Context(), c.Response())
		}
		return f(c)
	}
}
