package persistence

import (
	"context"
	"github.com/songfei1983/go-api-server/domain/model"
	"github.com/songfei1983/go-api-server/domain/repository"
)

type userPersistence struct {}

func NewUserPersistence() repository.UserRepository {
	return &userPersistence{}
}

func (u userPersistence) GetAll(ctx context.Context) ([]*model.User, error) {
	u1 := &model.User{"hoge", "abc"}
	u2 := &model.User{"hoge", "abc"}

	return []*model.User{u1, u2}, nil
}

