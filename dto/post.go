package dto

type PostData struct {
	Content  string `json:"content"`
	ImageUrl string `json:"image_url" validate:"required,url"`
}

type ViewPost struct {
	PostID   uint   `json:"post_id" extensions:"x-order=0"`
	Content  string `json:"content" extensions:"x-order=1"`
	ImageUrl string `json:"image_url" extensions:"x-order=2"`
	UserInfo `json:"user" extensions:"x-order=3"`
}

type ViewPostWithComments struct {
	ViewPost `json:"post" extensions:"x-order=0"`
	Comments []Comment `json:"comments" extensions:"x-order=1"`
}

type ViewPostWithoutUserInfo struct {
	PostID   uint   `json:"post_id" extensions:"x-order=0"`
	Content  string `json:"content" extensions:"x-order=1"`
	ImageUrl string `json:"image_url" extensions:"x-order=2"`
}
