package repository

import (
	"context"
	"github.com/songfei1983/go-api-server/internal/user/domain/model"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]*model.User, error)
	CreateUser(ctx context.Context) error
}
