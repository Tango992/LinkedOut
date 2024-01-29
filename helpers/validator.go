package helpers

import (
	"graded-3/utils"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	NewValidator *validator.Validate
}

// Custom validator using go-playground/validator
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.NewValidator.Struct(i); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.Details(err.Error()))
	}
	return nil
}
