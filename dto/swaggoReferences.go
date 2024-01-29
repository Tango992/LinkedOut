package dto

import "graded-3/entity"

type ActivitiesResponse struct {
	Message string `json:"message" extensions:"x-order=0"`
	Data    []Log  `json:"data" extensions:"x-order=1"`
}

type UserWithoutPassword struct {
	ID       uint   `gorm:"primaryKey" json:"id" extensions:"x-order=0"`
	FullName string `gorm:"not null" json:"full_name" extensions:"x-order=1"`
	Email    string `gorm:"not null;unique" json:"email" extensions:"x-order=2"`
	Username string `gorm:"not null;unique" json:"username" extensions:"x-order=3"`
	Age      uint   `gorm:"not null" json:"age" extensions:"x-order=5"`
}

type RegisterResponse struct {
	Message string              `json:"message" extensions:"x-order=0"`
	Data    UserWithoutPassword `json:"data" extensions:"x-order=1"`
}

type LoginResponse struct {
	Message string `json:"message" extensions:"x-order=0"`
	Data    string `json:"data" extensions:"x-order=1"`
}

type PostAndDeleteCommentResponse struct {
	Message string         `json:"message" extensions:"x-order=0"`
	Data    entity.Comment `json:"data" extensions:"x-order=1"`
}

type GetCommentResponse struct {
	Message string      `json:"message" extensions:"x-order=0"`
	Data    ViewComment `json:"data" extensions:"x-order=1"`
}

type PostAndDeleteResponse struct {
	Message string      `json:"message" extensions:"x-order=0"`
	Data    entity.Post `json:"data" extensions:"x-order=1"`
}

type GetPostByIdResponse struct {
	Message string               `json:"message" extensions:"x-order=0"`
	Data    ViewPostWithComments `json:"data" extensions:"x-order=1"`
}

type GetAllPostsResponse struct {
	Message string     `json:"message" extensions:"x-order=0"`
	Data    []ViewPost `json:"data" extensions:"x-order=1"`
}
