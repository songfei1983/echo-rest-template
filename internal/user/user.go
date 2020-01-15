package user

import (
	"github.com/songfei1983/go-api-server/internal/model"
)

type UseCase interface {
	GetAll() ([]*model.User, error)
}

func NewUseCase(r Repository) UseCase {
	return &useCase{repository: r}
}

type useCase struct {
	repository Repository
}

func (u useCase) GetAll() ([]*model.User, error) {
	return u.repository.GetAllUser()
}
