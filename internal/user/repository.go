package user

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/songfei1983/go-api-server/cmd/api/app"
	"github.com/songfei1983/go-api-server/helper"
	"github.com/songfei1983/go-api-server/internal/entity"
	"github.com/songfei1983/go-api-server/internal/model"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*model.User, error)
	CreateUser(ctx helper.CustomContext, m *model.CreateUser) error
	GetByID(id int) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
}

func NewUserPersistence(api *app.APP) Repository {
	return &userPersistence{api.DB}
}

type userPersistence struct {
	db *gorm.DB
}

func (u *userPersistence) GetByID(id int) (*model.User, error) {
	panic("implement me")
}

func (u *userPersistence) GetByEmail(email string) (*model.User, error) {
	e := &entity.User{}
	if err := u.db.Where("email = ?", email).First(e).Error; err != nil {
		return nil, err
	}
	// recovery for verify password
	m := toUserModel(e)
	m.Password = model.Password(e.Password)
	return m, nil
}

func (u *userPersistence) GetAll(ctx context.Context) ([]*model.User, error) {
	entities := make([]*entity.User, 0)
	err := u.db.Find(&entities).Error
	models := make([]*model.User, len(entities))
	for k, v := range entities {
		models[k] = toUserModel(v)
	}
	return models, err
}

func (u *userPersistence) CreateUser(ctx helper.CustomContext, m *model.CreateUser) error {
	return u.db.Save(fromUserModel(m)).Error
}

func toUserModel(e *entity.User) *model.User {
	u := new(model.User)
	u.ID = e.ID
	u.Name = e.Name
	u.Email = e.Email
	u.IsEnabled = e.IsEnabled
	u.Password = u.Password.Mask()
	return u
}
func fromUserModel(i interface{}) *entity.User {
	e := new(entity.User)
	switch m := i.(type) {
	case *model.CreateUser:
		e.Name = m.Name
		e.Email = m.Email
		e.Password = m.Password.HashAndSalt()
		e.IsEnabled = true
		e.Role = m.Role
	}
	return e
}
