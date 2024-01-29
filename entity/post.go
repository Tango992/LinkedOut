package entity

type Post struct {
	ID       uint      `gorm:"primaryKey" json:"post_id" extensions:"x-order=0"`
	UserID   uint      `gorm:"not null" json:"user_id" extensions:"x-order=1"`
	Content  string    `gorm:"not null" json:"content" extensions:"x-order=2"`
	ImageUrl string    `gorm:"not null" json:"image_url" extensions:"x-order=3"`
	Comments []Comment `json:"-"`
}