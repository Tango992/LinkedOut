package dto

type Comment struct {
	ID       uint   `json:"comment_id" extensions:"x-order=0"`
	Username string `json:"username" extensions:"x-order=1"`
	Comment  string `json:"comment" extensions:"x-order=2"`
}

type PostComment struct {
	Comment string `json:"comment" validate:"required" extensions:"x-order=0"`
	PostID  uint   `json:"post_id" validate:"required" extensions:"x-order=1"`
}

type ViewComment struct {
	CommentID               uint   `json:"comment_id" extensions:"x-order=0"`
	Comment                 string `json:"comment" extensions:"x-order=1"`
	UserInfo                `json:"user" extensions:"x-order=2"`
	ViewPostWithoutUserInfo `json:"post" extensions:"x-order=3"`
}
