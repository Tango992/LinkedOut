package controller

import (
	"graded-3/dto"
	"graded-3/entity"
	"graded-3/helpers"
	"graded-3/repository"
	"graded-3/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CommentController struct {
	repository.DbHandler
}

func NewCommentController(dbHandler repository.DbHandler) CommentController {
	return CommentController{
		DbHandler: dbHandler,
	}
}

// Comments      godoc
// @Summary      Post a new comment
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "JWT Token"
// @Param        request body dto.PostComment  true  "Comment data"
// @Success      201  {object}  dto.PostAndDeleteCommentResponse
// @Failure      400  {object}  utils.ErrResponse
// @Failure      401  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /comments [post]
func (cc CommentController) PostComment(c echo.Context) error {
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	var commentDataTmp dto.PostComment
	if err := c.Bind(&commentDataTmp); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.Details(err.Error()))
	}

	if err := c.Validate(&commentDataTmp); err != nil {
		return err
	}

	commentData := entity.Comment{
		UserID: user.ID,
		PostID: commentDataTmp.PostID,
		Comment: commentDataTmp.Comment,
	}
	
	if err := cc.DbHandler.AddCommentIntoDb(user, &commentData); err != nil {
		return err
	}
	
	return c.JSON(http.StatusCreated, entity.Response{
		Message: "Comment posted",
		Data: commentData,
	})
}

// Comments      godoc
// @Summary      Get comment by id
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Comment ID"
// @Param        Authorization header string true "JWT Token"
// @Success      200  {object}  dto.GetCommentResponse
// @Failure      400  {object}  utils.ErrResponse
// @Failure      401  {object}  utils.ErrResponse
// @Failure      404  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /comments/{id} [get]
func (cc CommentController) GetComment(c echo.Context) error {
	commentIdTmp := c.Param("id")
	commentId, err := strconv.Atoi(commentIdTmp)
	if err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.Details(err.Error()))
	}

	commentData, err := cc.DbHandler.GetCommentFromDb(uint(commentId))
	if err != nil {
		return err
	}
	
	return c.JSON(http.StatusOK, entity.Response{
		Message: "Get comment by ID",
		Data: commentData,
	})
}

// Comments      godoc
// @Summary      Delete comment by id
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Comment ID"
// @Param        Authorization header string true "JWT Token"
// @Success      200  {object}  dto.PostAndDeleteCommentResponse
// @Failure      400  {object}  utils.ErrResponse
// @Failure      401  {object}  utils.ErrResponse
// @Failure      404  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /comments/{id} [delete]
func (cc CommentController) DeleteComment(c echo.Context) error {
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}
	
	commentIdTmp := c.Param("id")
	commentId, err := strconv.Atoi(commentIdTmp)
	if err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.Details(err.Error()))
	}

	commentData := entity.Comment{
		ID: uint(commentId),
		UserID: user.ID,
	}

	if err := cc.DbHandler.DeleteCommentFromDb(user, &commentData); err != nil {
		return err
	}
	
	return c.JSON(http.StatusOK, entity.Response{
		Message: "Delete comment by ID",
		Data: commentData,
	})
}
