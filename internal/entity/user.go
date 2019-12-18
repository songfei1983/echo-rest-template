package entity

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name      string
	Role      string
	Email     string
	Password  string
	IsEnabled bool
}
