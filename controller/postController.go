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

type PostController struct {
	repository.DbHandler
}

func NewPostController(dbHandler repository.DbHandler) PostController {
	return PostController{
		DbHandler: dbHandler,
	}
}

// Post      godoc
// @Summary      Create a new post
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "JWT Token"
// @Param        request body dto.PostData  true  "Post data"
// @Success      201  {object}  dto.PostAndDeleteResponse
// @Failure      400  {object}  utils.ErrResponse
// @Failure      401  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /posts [post]
func (pc PostController) AddPost(c echo.Context) error {
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	var postDataTmp dto.PostData
	if err := c.Bind(&postDataTmp); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.Details(err.Error()))
	}

	if err := c.Validate(&postDataTmp); err != nil {
		return err
	}

	if err := helpers.CheckAndPopulateContent(&postDataTmp); err != nil {
		return err
	}

	postData := entity.Post{
		UserID:   user.ID,
		Content:  postDataTmp.Content,
		ImageUrl: postDataTmp.ImageUrl,
	}

	if err := pc.DbHandler.AddPostIntoDb(user, &postData); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, entity.Response{
		Message: "Posted",
		Data:    postData,
	})
}

// Posts      godoc
// @Summary      Get all posts
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "JWT Token"
// @Success      200  {object}  dto.GetAllPostsResponse
// @Failure      400  {object}  utils.ErrResponse
// @Failure      401  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /posts [get]
func (pc PostController) GetAllPosts(c echo.Context) error {
	posts, err := pc.DbHandler.FindAllPostInDb()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, entity.Response{
		Message: "Get all posts",
		Data:    posts,
	})
}

// Posts      godoc
// @Summary      Get post by id
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Post ID"
// @Param        Authorization header string true "JWT Token"
// @Success      200  {object}  dto.GetPostByIdResponse
// @Failure      400  {object}  utils.ErrResponse
// @Failure      401  {object}  utils.ErrResponse
// @Failure      404  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /posts/{id} [get]
func (pc PostController) GetPostByID(c echo.Context) error {
	postIdTmp := c.Param("id")
	postId, err := strconv.Atoi(postIdTmp)
	if err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.Details(err.Error()))
	}

	postData, err := pc.DbHandler.FindDetailedPostInDb(uint(postId))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, entity.Response{
		Message: "Get post by ID",
		Data:    postData,
	})
}

// Posts      godoc
// @Summary      Delete post by id
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Post ID"
// @Param        Authorization header string true "JWT Token"
// @Success      200  {object}  dto.PostAndDeleteResponse
// @Failure      400  {object}  utils.ErrResponse
// @Failure      401  {object}  utils.ErrResponse
// @Failure      404  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /posts/{id} [delete]
func (pc PostController) DeletePostById(c echo.Context) error {
	postIdTmp := c.Param("id")
	postId, err := strconv.Atoi(postIdTmp)
	if err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.Details(err.Error()))
	}

	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	postData := entity.Post{
		ID:     uint(postId),
		UserID: user.ID,
	}

	if err := pc.DbHandler.DeletePostFromDb(user, &postData); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, entity.Response{
		Message: "Delete post by ID",
		Data:    postData,
	})
}
