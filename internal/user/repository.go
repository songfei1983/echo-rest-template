package user

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	api "github.com/songfei1983/go-api-server/cmd/api/app"
	"github.com/songfei1983/go-api-server/internal/entity"
	"github.com/songfei1983/go-api-server/internal/model"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*model.User, error)
	CreateUser(ctx context.Context) error
}

type userPersistence struct {
	db *gorm.DB
}

func NewUserPersistence(app *api.APP) Repository {
	return &userPersistence{app.DB}
}

func (u userPersistence) GetAll(ctx context.Context) ([]*model.User, error) {
	entities := make([]*entity.User, 0)
	err := u.db.Find(&entities).Error
	models := make([]*model.User, len(entities))
	for k, v := range entities {
		models[k] = toUserModel(v)
	}
	return models, err
}

func (u userPersistence) CreateUser(ctx context.Context) error {
	c := ctx.Value("ctx").(echo.Context)
	m := c.Get("data").(model.User)
	return u.db.Save(fromUserModel(&m)).Error
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
func fromUserModel(m *model.User) *entity.User {
	e := new(entity.User)
	e.ID = m.ID
	e.Name = m.Name
	e.Email = m.Email
	e.Password = m.Password.HashAndSalt()
	e.IsEnabled = m.IsEnabled
	e.Role = m.Role
	return e
}
