package user

import (
	"github.com/songfei1983/go-api-server/cmd/api/app"
	"github.com/songfei1983/go-api-server/internal/entity"
	"github.com/songfei1983/go-api-server/internal/model"
	"github.com/songfei1983/go-api-server/internal/persistence"
)

type Repository interface {
	GetAllUser() ([]*entity.User, error)
	CreateUser(m *entity.User) error
	GetUserByID(id int) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
}

func NewUserRepository(api *app.APP) Repository {
	return &userRepository{persistence.Persistence{DB: api.DB}}
}

type userRepository struct {
	persistence.Persistence
}

func ToUserModel(e *entity.User) *model.User {
	u := new(model.User)
	u.ID = e.ID
	u.Name = e.Name
	u.Email = e.Email
	u.IsEnabled = e.IsEnabled
	u.Password = u.Password.Mask()
	return u
}
func FromUserModel(i interface{}) *entity.User {
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
