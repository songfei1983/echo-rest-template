package model

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// Data
type UserModel struct {
	Name  string `json:"name"`
	Role  string `json:"role"`
	Email string `json:"email"`
}

// Response
type User struct {
	ID        uint     `json:"id"`
	Password  Password `json:"password"`
	IsEnabled bool     `json:"is_enabled"`
	UserModel
}

// Value
type Password string

func (p Password) String() string {
	return string(p)
}
func (p Password) HashAndSalt() string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	return string(hash)
}

func (p Password) Mask() Password {
	return "******"
}

func (p Password) Verify(plainPwd Password) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p), []byte(plainPwd))
	fmt.Println(err)
	return err == nil
}

// Request
type CreateUser struct {
	UserModel
	Password Password `json:"password"`
}

type EditUser struct {
	ID        uint `json:"id"`
	IsEnabled bool `json:"is_enabled"`
	UserModel
}

type ChangePassword struct {
	ID          uint     `json:"id"`
	Password    Password `json:"password"`
	NewPassword Password `json:"newPassword"`
}
