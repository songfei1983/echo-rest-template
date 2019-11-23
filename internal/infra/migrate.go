package infra

import (
	"github.com/jinzhu/gorm"
	"github.com/songfei1983/go-api-server/internal/infra/entity"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&entity.User{})
}
