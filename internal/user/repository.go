package user

import (
	"github.com/jinzhu/gorm"
	"github.com/songfei1983/go-api-server/cmd/api/app"
	"github.com/songfei1983/go-api-server/internal/model"
)

type Repository interface {
	GetAllUser() ([]*model.User, error)
}

func NewUserRepository(api *app.APP) Repository {
	return &repository{api.DB}
}

type repository struct {
	DB *gorm.DB
}

func (r repository) GetAllUser() (users []*model.User, err error) {
	err = r.DB.Find(&users).Error
	return users, err
}
