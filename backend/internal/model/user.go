package model

type User struct {
	Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
