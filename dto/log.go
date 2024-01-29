package dto

type Log struct {
	Description string `json:"description" extensions:"x-order=0"`
	CreatedAt   string `json:"created_at" extensions:"x-order=1"`
}
