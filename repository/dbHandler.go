package repository

import (
	"errors"
	"fmt"
	"graded-3/dto"
	"graded-3/entity"
	"graded-3/utils"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DbHandler struct {
	*gorm.DB
}

func NewDbHandler(db *gorm.DB) DbHandler {
	return DbHandler{
		DB: db,
	}
}

func (db DbHandler) AddUserIntoDb(user *entity.User) error {
	if err := db.Create(user).Error; err != nil {
		return echo.NewHTTPError(utils.ErrConflict.Details(err.Error()))
	}
	return nil
}

func (db DbHandler) FindUserInDb(loginData dto.Login) (entity.User, error) {
	var user entity.User

	res := db.Where("email = ?", loginData.Email).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return entity.User{}, echo.NewHTTPError(utils.ErrNotFound.Details(res.Error.Error()))
		}
		return entity.User{}, echo.NewHTTPError(utils.ErrInternalServer.Details(res.Error.Error()))
	}
	return user, nil
}

func (db DbHandler) AddPostIntoDb(user entity.Claims, data *entity.Post) error {
	txErr := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(data).Error; err != nil {
			return err
		}

		time := time.Now().Format("2006-01-02 15:04:05")
		description := fmt.Sprintf("@%v created new POST with ID %v", user.Username, data.ID)
		logData := entity.UserActivityLog{
			UserID:      user.ID,
			Description: description,
			CreatedAt:   time,
		}

		if err := tx.Create(&logData).Error; err != nil {
			return err
		}
		return nil
	})
	if txErr != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.Details(txErr.Error()))
	}
	return nil
}

func (db DbHandler) FindAllPostInDb() ([]dto.ViewPost, error) {
	var posts []dto.ViewPost

	if err := db.Table("posts p").Select("p.id AS post_id, u.id AS user_id, u.username, u.email, p.content, p.image_url").Joins("JOIN users u ON u.id = p.user_id").Scan(&posts).Error; err != nil {
		return []dto.ViewPost{}, echo.NewHTTPError(utils.ErrInternalServer.Details(err.Error()))
	}
	return posts, nil
}

func (db DbHandler) FindDetailedPostInDb(postID uint) (dto.ViewPostWithComments, error) {
	var (
		postData = dto.ViewPost{}
		comments = []dto.Comment{}
	)

	txErr := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("posts p").Select("p.id AS post_id, u.id AS user_id, u.username, u.email, p.content, p.image_url").Joins("JOIN users u ON u.id = p.user_id").Where("p.id = ?", postID).Take(&postData).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return echo.NewHTTPError(utils.ErrNotFound.Details("Post does not exist"))
			}
			return echo.NewHTTPError(utils.ErrInternalServer.Details(err.Error()))
		}

		if err := tx.Table("comments c").Select("c.id, u.username, c.comment").Joins("JOIN posts p ON p.id = c.post_id").Joins("JOIN users u ON u.id = c.user_id").Where("p.id = ?", postID).Order("c.id").Scan(&comments).Error; err != nil {
			return echo.NewHTTPError(utils.ErrInternalServer.Details(err.Error()))
		}
		return nil
	})
	if txErr != nil {
		return dto.ViewPostWithComments{}, txErr
	}

	combinedData := dto.ViewPostWithComments{
		ViewPost: postData,
		Comments: comments,
	}
	return combinedData, nil
}

func (db DbHandler) DeletePostFromDb(user entity.Claims, postData *entity.Post) error {
	var post entity.Post

	txErr := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", postData.ID).First(&post).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return echo.NewHTTPError(utils.ErrNotFound.Details("Post does not exist"))
			}
			return echo.NewHTTPError(utils.ErrInternalServer.Details(err.Error()))
		}

		res := tx.Select(clause.Associations).Where("user_id = ?", postData.UserID).Delete(postData)
		if res.Error != nil {
			return echo.NewHTTPError(utils.ErrInternalServer.Details(res.Error.Error()))
		}

		if res.RowsAffected == 0 {
			return echo.NewHTTPError(utils.ErrUnauthorized.Details("You don't have permission to delete other user's post"))
		}

		time := time.Now().Format("2006-01-02 15:04:05")
		description := fmt.Sprintf("@%v deleted a POST with ID %v", user.Username, postData.ID)
		logData := entity.UserActivityLog{
			UserID:      user.ID,
			Description: description,
			CreatedAt:   time,
		}

		if err := tx.Create(&logData).Error; err != nil {
			return err
		}
		return nil
	})
	if txErr != nil {
		return txErr
	}

	postData.Content = post.Content
	postData.ImageUrl = post.ImageUrl
	return nil
}

func (db DbHandler) AddCommentIntoDb(user entity.Claims, data *entity.Comment) error {
	txErr := db.Transaction(func(tx *gorm.DB) error {
		var exists bool
		if db.Raw("SELECT EXISTS(SELECT * FROM posts p WHERE p.id = ?)", data.PostID).Take(&exists); !exists {
			return echo.NewHTTPError(utils.ErrNotFound.Details("Post does not exist"))
		}

		if err := db.Create(data).Error; err != nil {
			return echo.NewHTTPError(utils.ErrInternalServer.Details(err.Error()))
		}

		time := time.Now().Format("2006-01-02 15:04:05")
		description := fmt.Sprintf("@%v created new COMMENT on Comment ID %v", user.Username, data.ID)
		logData := entity.UserActivityLog{
			UserID:      user.ID,
			Description: description,
			CreatedAt:   time,
		}

		if err := tx.Create(&logData).Error; err != nil {
			return err
		}

		return nil
	})
	if txErr != nil {
		return txErr
	}
	return nil
}

func (db DbHandler) GetCommentFromDb(commentId uint) (dto.ViewComment, error) {
	var combinedData dto.ViewComment

	if err := db.Table("comments c").Select("c.id AS comment_id, c.comment AS comment, u.id AS user_id, u.email, u.username, p.id AS post_id, p.content, p.image_url").Joins("JOIN posts p ON c.post_id = p.id").Joins("JOIN users u ON c.user_id = u.id").Where("c.id = ?", commentId).Take(&combinedData).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.ViewComment{}, echo.NewHTTPError(utils.ErrNotFound.Details("Comment does not exist"))
		}
		return dto.ViewComment{}, echo.NewHTTPError(utils.ErrInternalServer.Details(err.Error()))
	}
	return combinedData, nil
}

func (db DbHandler) DeleteCommentFromDb(user entity.Claims, commentData *entity.Comment) error {
	var comment entity.Comment
	
	txErr := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", commentData.ID).First(&comment).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return echo.NewHTTPError(utils.ErrNotFound.Details("Comment does not exist"))
			}
			return echo.NewHTTPError(utils.ErrInternalServer.Details(err.Error()))
		}
		
		res := tx.Select(clause.Associations).Where("user_id = ?", user.ID).Delete(commentData)
		if res.Error != nil {
			return echo.NewHTTPError(utils.ErrInternalServer.Details(res.Error.Error()))
		}
		
		if res.RowsAffected == 0 {
			return echo.NewHTTPError(utils.ErrUnauthorized.Details("You don't have permission to delete other user's comment"))
		}

		time := time.Now().Format("2006-01-02 15:04:05")
		description := fmt.Sprintf("@%v deleted a COMMENT on Comment ID %v", user.Username, comment.ID)
		logData := entity.UserActivityLog{
			UserID:      user.ID,
			Description: description,
			CreatedAt:   time,
		}

		if err := tx.Create(&logData).Error; err != nil {
			return err
		}
		return nil
	})
	if txErr != nil {
		return txErr
	}
	
	commentData.ID = comment.ID
	commentData.UserID = comment.UserID
	commentData.PostID = comment.PostID
	commentData.Comment = comment.Comment
	return nil
}

func (db DbHandler) FindUserLog(userId uint) ([]dto.Log, error) {
	logs := []dto.Log{}

	if err := db.Table("user_activity_logs").Select("description", "created_at").Where("user_id = ?", userId).Scan(&logs).Error; err != nil {
		return []dto.Log{}, echo.NewHTTPError(utils.ErrInternalServer.Details(err.Error()))
	}
	
	return logs, nil
}