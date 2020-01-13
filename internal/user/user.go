package user

import (
	"context"
	"github.com/songfei1983/go-api-server/internal/model"
	"github.com/songfei1983/go-api-server/pkg/helper"
	"github.com/songfei1983/go-api-server/pkg/logger"
)

type UseCase interface {
	GetAll(ctx context.Context) ([]*model.User, error)
	Create(ctx helper.CustomContext, m *model.CreateUser) error
	Update(ctx context.Context) error
	Delete(ctx context.Context) error
	GetByID(ctx context.Context) (*model.User, error)
	GetByEmail(ctx helper.CustomContext, email string) (*model.User, error)
}

func NewUseCase(r Repository) UseCase {
	return &useCase{
		userRepository: r,
	}
}

type useCase struct {
	userRepository Repository
}

func (u *useCase) Create(ctx helper.CustomContext, m *model.CreateUser) error {
	e := FromUserModel(m)
	if err := u.userRepository.CreateUser(e); err != nil {
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
	res, err := u.userRepository.GetAllUser()
	if err != nil {
		return nil, err
	}
	ms := make([]*model.User, len(res))
	for k,v := range res {
		ms[k] = ToUserModel(v)
	}
	return ms, nil
}

func (u *useCase) GetByEmail(cc helper.CustomContext, email string) (*model.User, error) {
	e, err := u.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return ToUserModel(e), err
}

