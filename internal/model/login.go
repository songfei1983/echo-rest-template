package model

type LoginUser struct {
	Email    string   `json:"email"`
	Password Password `json:"password"`
}
