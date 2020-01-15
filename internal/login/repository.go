package login

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/songfei1983/go-api-server/cmd/api/app"
	"github.com/songfei1983/go-api-server/internal/model"
)

type Repository interface {
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(user *model.User) error
}

type MySQLRepository struct {
	DB     *gorm.DB
	Logger echo.Logger
}

func (m MySQLRepository) CreateUser(user *model.User) error {
	pw := model.Password(user.Password)
	user.Password = pw.HashAndSalt()
	return m.DB.Save(user).Error
}

func (m MySQLRepository) GetUserByEmail(email string) (*model.User, error) {
	entity := new(model.User)
	err := m.DB.Where("email = ?", email).First(entity).Error
	return entity, err
}

func NewRepository(api *app.APP) Repository {
	return &MySQLRepository{DB: api.DB, Logger: api.Server.Logger}
}
