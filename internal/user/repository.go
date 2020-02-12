package user

import (
	"github.com/jinzhu/gorm"
	"github.com/songfei1983/go-api-server/internal/server"
	"github.com/songfei1983/go-api-server/internal/model"
)

type Repository interface {
	GetAllUser() ([]*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(user *model.User) error
}

func NewUserRepository(api *server.API) Repository {
	return &repository{api.DB}
}

type repository struct {
	DB *gorm.DB
}

func (r *repository) GetAllUser() (users []*model.User, err error) {
	err = r.DB.Find(&users).Error
	return users, err
}

func (r *repository) CreateUser(user *model.User) error {
	pw := model.Password(user.Password)
	user.Password = pw.HashAndSalt()
	return r.DB.Save(user).Error
}

func (r *repository) GetUserByEmail(email string) (*model.User, error) {
	entity := new(model.User)
	err := r.DB.Where("email = ?", email).First(entity).Error
	return entity, err
}
