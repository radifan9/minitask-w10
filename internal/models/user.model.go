package models

type User struct {
	Id       string `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}
