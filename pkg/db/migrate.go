package db

import (
	"github.com/jinzhu/gorm"
	"github.com/songfei1983/go-api-server/internal/model"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}
