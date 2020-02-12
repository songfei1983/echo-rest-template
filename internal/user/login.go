package user

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/songfei1983/go-api-server/internal/model"
	"github.com/songfei1983/go-api-server/internal/server"
	"github.com/songfei1983/go-api-server/pkg/helper"
	"net/http"
)

type LoginForm struct {
	Email    string `validate:"required,gt=0"`
	Password string `validate:"required,gt=0"`
}

func NewLogin(a *server.API) server.Service {
	r := NewUserRepository(a)
	return &LoginService{repo: r}
}

var _ server.Service = (*LoginService)(nil)

type LoginService struct {
	repo Repository
}

func (u *LoginService) Execute(c context.Context) (*server.Response, error) {
	data := c.Value("data")
	form, ok := data.(*LoginForm)
	if !ok {
		return nil, echo.ErrBadRequest
	}
	e, err := u.repo.GetUserByEmail(form.Email)
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
		return &server.Response{
			StatusCode: http.StatusOK,
			Data:       map[string]string{"token": t},
		}, nil
	}
	return nil, echo.ErrUnauthorized
}
func (u *LoginService) RequestSchema() interface{} {
	req := new(LoginForm)
	return req
}
