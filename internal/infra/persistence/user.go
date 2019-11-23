package persistence

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"

	api "github.com/songfei1983/go-api-server/cmd/api/app"
	"github.com/songfei1983/go-api-server/internal/infra/entity"
	"github.com/songfei1983/go-api-server/internal/user/domain/model"
	"github.com/songfei1983/go-api-server/internal/user/domain/repository"
)

type userPersistence struct {
	db *gorm.DB
}

func NewUserPersistence(app *api.APP) repository.UserRepository {
	return &userPersistence{app.DB}
}

func (u userPersistence) GetAll(ctx context.Context) ([]*model.User, error) {
	entities := make([]*entity.User, 0)
	err := u.db.Find(&entities).Error
	models := make([]*model.User, len(entities))
	for k, v := range entities {
		models[k] = transform(v)
	}
	return models, err
}

func (u userPersistence) CreateUser(ctx context.Context) error {
	c := ctx.Value("ctx").(echo.Context)
	m := c.Get("data").(model.User)
	return u.db.Save(&m).Error
}

func transform(e *entity.User) *model.User {
	return &model.User{
		ID:   e.ID,
		Name: e.Name,
		Age:  e.Age,
	}
}
