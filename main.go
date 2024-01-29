package main

import (
	"graded-3/config"
	"graded-3/controller"
	"graded-3/helpers"
	"graded-3/middlewares"
	"graded-3/repository"
	_ "graded-3/docs"

	"github.com/go-playground/validator"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger" 
)

// @title Graded Challenge 3 API
// @version 1.0
// @description Made for Graded Challenge 3 - Hacktiv8 FTGO

// @contact.name Daniel Osvaldo Rahmanto
// @contact.email email@mail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	db := config.InitDb()
	dbHandler := repository.NewDbHandler(db)
	userController := controller.NewUserController(dbHandler)
	postController := controller.NewPostController(dbHandler)
	commentController := controller.NewCommentController(dbHandler)
	activityController := controller.NewActivityController(dbHandler)

	e := echo.New()
	e.Validator = &helpers.CustomValidator{NewValidator: validator.New()}

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.RequestLoggerWithConfig(config.Zap()))
	e.Use(middleware.Recover())
	
	
	users := e.Group("/users")
	{
		users.POST("/register", userController.Register)
		users.POST("/login", userController.Login)
	}
	
	posts := e.Group("/posts")
	posts.Use(middlewares.RequireAuth)
	{
		posts.POST("", postController.AddPost)
		posts.GET("", postController.GetAllPosts)
		posts.GET("/:id", postController.GetPostByID)
		posts.DELETE("/:id", postController.DeletePostById)
	}
	
	comments := e.Group("/comments")
	comments.Use(middlewares.RequireAuth)
	{
		comments.POST("", commentController.PostComment)
		comments.GET("/:id", commentController.GetComment)
		comments.DELETE("/:id", commentController.DeleteComment)
	}
	
	e.GET("/activities", middlewares.RequireAuth(activityController.GetActivities))
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
