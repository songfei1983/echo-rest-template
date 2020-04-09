package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) (*Repository, error) {
	if db == nil {
		return nil, fmt.Errorf("no connection")
	}
	return &Repository{db: db}, nil
}
