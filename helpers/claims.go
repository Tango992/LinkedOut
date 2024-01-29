package helpers

import (
	"graded-3/entity"
	"graded-3/utils"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// Get claims from Echo Context and assert the corresponding data types
func GetClaims(c echo.Context) (entity.Claims, error) {
	claimsTmp := c.Get("user")
	if claimsTmp == nil {
		return entity.Claims{}, echo.NewHTTPError(utils.ErrUnauthorized.Details("Failed to fetch user claims from JWT"))
	}

	claims := claimsTmp.(jwt.MapClaims)
	return entity.Claims{
		ID: uint(claims["id"].(float64)),
		Email: claims["email"].(string),
		Username: claims["username"].(string),
		FullName: claims["full_name"].(string),
	}, nil
}