package dto

type Register struct {
	FullName string `json:"full_name" validate:"required" extensions:"x-order=0"`
	Email    string `json:"email" validate:"required,email" extensions:"x-order=1"`
	Username string `json:"username" validate:"required" extensions:"x-order=2"`
	Password string `json:"password" validate:"required" extensions:"x-order=3"`
	Age      uint   `json:"age" validate:"required,min=12" extensions:"x-order=4"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email" extensions:"x-order=0"`
	Password string `json:"password" validate:"required" extensions:"x-order=1"`
}

type UserInfo struct {
	UserID   uint   `json:"user_id" extensions:"x-order=0"`
	Username string `json:"username" extensions:"x-order=1"`
	Email    string `json:"email" extensions:"x-order=2"`
}
