package helpers

import (
	"graded-3/entity"
	"graded-3/utils"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// Create a new JWT and sign it
func SignNewJWT(user entity.User) (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(4 * time.Hour).Unix(),
		"id": user.ID,
		"email": user.Email,
		"username": user.Username,
		"full_name": user.FullName,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", echo.NewHTTPError(utils.ErrInternalServer.Details(err.Error()))
	}
	return tokenString, nil
}