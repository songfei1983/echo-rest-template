package persistence

import (
	"github.com/songfei1983/go-api-server/internal/entity"
)

func (r *Persistence) GetUserByEmail(email string) (*entity.User, error) {
	e := &entity.User{}
	if err := r.DB.Where("email = ?", email).First(e).Error; err != nil {
		return nil, err
	}
	return e, nil
}

func (r *Persistence) GetAllUser() ([]*entity.User, error) {
	entities := make([]*entity.User, 0)
	err := r.DB.Find(&entities).Error
	return entities, err
}

func (r *Persistence) CreateUser(m *entity.User) error {
	return r.DB.Save(m).Error
}

func (r *Persistence) GetUserByID(id int) (*entity.User, error) {
	e := new(entity.User)
	if err := r.DB.First(e, id).Error; err != nil {
		return nil, err
	}
	return e, nil
}
