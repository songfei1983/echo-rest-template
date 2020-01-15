package model

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name      string `validate:"required"`
	Role      string
	Email     string `validate:"required"`
	Password  string `validate:"required"`
	IsEnabled bool
}

func (u *User) MarshalJSON() ([]byte, error) {
	u.Password = Password(u.Password).Mask().String()
	return json.Marshal(*u)
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

func (p Password) Verify(plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p), []byte(plainPwd))
	fmt.Println(err)
	return err == nil
}
