package user

import (
	"context"
	"github.com/songfei1983/go-api-server/helper"
	"github.com/songfei1983/go-api-server/internal/model"
	"github.com/songfei1983/go-api-server/logger"
)

type UseCase interface {
	GetAll(ctx context.Context) ([]*model.User, error)
	Create(ctx context.Context) error
	Update(ctx context.Context) error
	Delete(ctx context.Context) error
	GetByID(ctx context.Context) (*model.User, error)
}

func NewUseCase(r Repository) UseCase {
	return &useCase{
		userRepository: r,
	}
}

type useCase struct {
	userRepository Repository
}

func (u *useCase) Create(ctx context.Context) error {
	if err := u.userRepository.CreateUser(ctx); err != nil {
		return err
	}
	return nil
}

func (u *useCase) Update(ctx context.Context) error {
	panic("implement me")
}

func (u *useCase) Delete(ctx context.Context) error {
	panic("implement me")
}

func (u *useCase) GetByID(ctx context.Context) (*model.User, error) {
	panic("implement me")
}

func (u *useCase) GetAll(ctx context.Context) ([]*model.User, error) {
	c := ctx.Value("ctx").(*helper.CustomContext)
	logger.Info(c, "Account name from context: %v", c.Get("AccountName"))
	res, err := u.userRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
