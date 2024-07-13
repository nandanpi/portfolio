package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

var JWTKey = []byte(os.Getenv("JWT_SECRET"))

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
		tokenString := c.Request().Header.Get("Authorization")
		if len(tokenString) == 0 {
			return c.JSON(401, c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"}))
		}
		tokenString = tokenString[7:]
		_, err := verifyToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
		}
		return f(c)
	}
}