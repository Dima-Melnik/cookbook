package model

type CookBook struct {
	Model
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
}

type CreateCookBook struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	UserID      uint   `json:"user_id"`
}

type UpdateCookBook struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
}
