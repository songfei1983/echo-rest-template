package usecase

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/songfei1983/go-api-server/internal/user/domain/model"
	"github.com/songfei1983/go-api-server/internal/user/domain/repository"
	"github.com/songfei1983/go-api-server/logger"
)

type UserUseCase interface {
	GetAll(ctx context.Context)([]*model.User, error)
	CreateUser(ctx context.Context) error
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository:r,
	}
}

func (u userUseCase) GetAll(ctx context.Context) ([]*model.User, error) {
	c := ctx.Value("ctx").(echo.Context)
	logger.Info(c, "Account name from context: %v", c.Get("AccountName"))
	res, err := u.userRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u userUseCase) CreateUser(ctx context.Context) error {
	if err := u.userRepository.CreateUser(ctx); err != nil {
		return err
	}
	return nil
}

