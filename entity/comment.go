package entity

type Comment struct {
	ID      uint   `gorm:"primary key" json:"comment_id" extensions:"x-order=0"`
	UserID  uint   `gorm:"not null" json:"user_id" extensions:"x-order=1"`
	PostID  uint   `gorm:"not null" json:"post_id" extensions:"x-order=2"`
	Comment string `gorm:"not null" json:"comment" extensions:"x-order=3"`
}
