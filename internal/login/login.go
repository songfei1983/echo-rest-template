package login

import (
	"context"
)

type UseCase interface {
	Login(c context.Context) error
	Logout(c context.Context) error
	Register(c context.Context) error
}

type useCase struct {}

func (u useCase) Login(c context.Context) error {
	panic("implement me")
}

func (u useCase) Logout(c context.Context) error {
	panic("implement me")
}

func (u useCase) Register(c context.Context) error {
	panic("implement me")
}

func NewUseCase() UseCase {
	return &useCase{}
}
