package entity

type User struct {
	ID               uint              `gorm:"primaryKey" json:"id" extensions:"x-order=0"`
	FullName         string            `gorm:"not null" json:"full_name" extensions:"x-order=1"`
	Email            string            `gorm:"not null;unique" json:"email" extensions:"x-order=2"`
	Username         string            `gorm:"not null;unique" json:"username" extensions:"x-order=3"`
	Password         string            `gorm:"not null" json:"password,omitempty" extensions:"x-order=4"`
	Age              uint              `gorm:"not null" json:"age" extensions:"x-order=5"`
	UserActivityLogs []UserActivityLog `json:"-"`
	Posts            []Post            `json:"-"`
	Comments         []Comment         `json:"-"`
}
