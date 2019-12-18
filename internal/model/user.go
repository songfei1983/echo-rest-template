package model

import "golang.org/x/crypto/bcrypt"

// Response
type User struct {
	ID        uint     `json:"id"`
	Name      string   `json:"name"`
	Role      string   `json:"role"`
	Email     string   `json:"email"`
	Password  Password `json:"password"`
	IsEnabled bool     `json:"is_enabled"`
}

// Value
type Password string

func (p Password) HashAndSalt() string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	return string(hash)
}

func (p Password) Mask() Password {
	return "******"
}

func (p Password) Verify(plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p), []byte(plainPwd))
	return err == nil
}

// Request
type CreateUser struct {
	Name     string   `json:"name"`
	Role     string   `json:"role"`
	Email    string   `json:"email"`
	Password Password `json:"password"`
}

type EditUser struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	IsEnabled bool   `json:"is_enabled"`
}

type ResetPassword struct {
	ID          uint     `json:"id"`
	OldPassword Password `json:"OldPassword"`
	Password    Password `json:"password"`
}

type LoginUser struct {
	Email    string   `json:"email"`
	Password Password `json:"password"`
}
