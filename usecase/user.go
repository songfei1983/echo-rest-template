package usecase

import (
	"context"
	"github.com/songfei1983/go-api-server/domain/model"
	"github.com/songfei1983/go-api-server/domain/repository"
)

type UserUseCase interface {
	GetAll(ctx context.Context)([]*model.User, error)
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
	res, err := u.userRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

