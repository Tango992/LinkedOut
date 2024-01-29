package controller

import (
	"graded-3/dto"
	"graded-3/entity"
	"graded-3/helpers"
	"graded-3/repository"
	"graded-3/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	repository.DbHandler
}

func NewUserController(dbHandler repository.DbHandler) UserController {
	return UserController{
		DbHandler: dbHandler,
	}
}

// Register      godoc
// @Summary      Register new user into database
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request body dto.Register  true  "Register data"
// @Success      201  {object}  dto.RegisterResponse
// @Failure      400  {object}  utils.ErrResponse
// @Failure      409  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /users/register [post]
func (uc UserController) Register(c echo.Context) error {
	var registerDataTmp dto.Register

	if err := c.Bind(&registerDataTmp); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.Details(err.Error()))
	}
	
	if err := c.Validate(&registerDataTmp); err != nil {
		return err
	}

	if err := helpers.CreateHash(&registerDataTmp); err != nil {
		return err
	}

	registerData := entity.User{
		FullName: registerDataTmp.FullName,
		Username: registerDataTmp.Username,
		Email: registerDataTmp.Email,
		Password: registerDataTmp.Password,
		Age: registerDataTmp.Age,
	}
	
	if err := uc.DbHandler.AddUserIntoDb(&registerData); err != nil {
		return err
	}

	registerData.Password = ""
	return c.JSON(http.StatusCreated, entity.Response{
		Message: "Registered",
		Data: registerData,
	})
}

// Login         godoc
// @Summary      Login existing user 
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request body dto.Login  true  "Login data"
// @Success      200  {object}  dto.LoginResponse
// @Failure      400  {object}  utils.ErrResponse
// @Failure      401  {object}  utils.ErrResponse
// @Failure      404  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /users/login [post]
func (uc UserController) Login(c echo.Context) error {
	var loginData dto.Login

	if err := c.Bind(&loginData); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.Details(err.Error()))
	}

	if err := c.Validate(&loginData); err != nil {
		return err
	}

	userData, err := uc.DbHandler.FindUserInDb(loginData)
	if err != nil {
		return err
	}

	if helpers.PasswordMismatch(userData, loginData) {
		return echo.NewHTTPError(utils.ErrUnauthorized.Details("Invalid username / password"))
	}

	jwtToken, err := helpers.SignNewJWT(userData)
	if err != nil {
		return err
	}
	
	return c.JSON(http.StatusOK, entity.Response{
		Message: "Logged in. Use the following token for authorization",
		Data: jwtToken,
	})
}