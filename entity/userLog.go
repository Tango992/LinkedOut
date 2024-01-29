package entity

type UserActivityLog struct {
	ID          uint   `gorm:"primaryKey" json:"id" extensions:"x-order=0"`
	UserID      uint   `gorm:"not null" json:"user_id" extensions:"x-order=1"`
	Description string `gorm:"not null" json:"description" extensions:"x-order=2"`
	CreatedAt   string `gorm:"not null;type:TIMESTAMP;default:NOW()" json:"created_at" extensions:"x-order=3"`
}
