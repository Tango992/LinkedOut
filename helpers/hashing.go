package helpers

import (
	"graded-3/dto"
	"graded-3/entity"
	"graded-3/utils"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// Create a new password hash
func CreateHash(data *dto.Register) error {
	hashed, err:= bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.Details(err.Error()))
	}
	data.Password = string(hashed)
	return nil
}

// Returns true if password does not match with the original hash
func PasswordMismatch(dbData entity.User, data dto.Login) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(dbData.Password), []byte(data.Password)); err != nil {
		return true
	}
	return false
}