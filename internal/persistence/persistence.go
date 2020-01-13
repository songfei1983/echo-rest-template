package persistence

import (
	"github.com/jinzhu/gorm"
)

type Persistence struct {
	DB *gorm.DB
}

