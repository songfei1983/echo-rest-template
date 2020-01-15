package login

import (
	"github.com/labstack/echo/v4"
	"github.com/songfei1983/go-api-server/internal/model"
	"github.com/songfei1983/go-api-server/pkg/helper"
)

type UseCase interface {
	Login(form *LoginForm) (token *string, err error)
	Register(user *model.User) error
}

type useCase struct {
	Repository Repository
	Logger     echo.Logger
}

func (u useCase) Register(user *model.User) error {
	return u.Repository.CreateUser(user)
}

func (u useCase) Login(form *LoginForm) (token *string, err error){
	e, err := u.Repository.GetUserByEmail(form.Email)
	if err != nil {
		return nil, err
	}
	var pw model.Password
	pw = model.Password(e.Password)
	if pw.Verify(form.Password) {
		t, err := helper.GenToken(*e)
		if err != nil {
			return nil, err
		}
		return &t, nil
	}
	return nil, echo.ErrUnauthorized
}

func NewUseCase(r Repository, l echo.Logger) UseCase {
	return &useCase{
		Repository: r,
		Logger:     l,
	}
}
