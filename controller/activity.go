package controller

import (
	"graded-3/entity"
	"graded-3/helpers"
	"graded-3/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ActivityController struct {
	repository.DbHandler
}

func NewActivityController(dbHandler repository.DbHandler) ActivityController {
	return ActivityController{
		DbHandler: dbHandler,
	}
}

// Activities    godoc
// @Summary      Get current user logs
// @Tags         activities
// @Produce      json
// @Param        Authorization header string true "JWT Token"
// @Success      200  {object}  dto.ActivitiesResponse
// @Failure      401  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /activities [get]
func (ac ActivityController) GetActivities(c echo.Context) error {
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}
	
	logs, err := ac.DbHandler.FindUserLog(user.ID)
	if err != nil {
		return err
	}
	
	return c.JSON(http.StatusOK, entity.Response{
		Message: "Get activities",
		Data: logs,
	})
}