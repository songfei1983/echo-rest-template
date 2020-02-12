package user

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/songfei1983/go-api-server/internal/model"
	"github.com/songfei1983/go-api-server/internal/server"
	"net/http"
)

func NewRegister(a *server.API) server.Service {
	r := NewUserRepository(a)
	return &RegisterService{repo: r}
}

var _ server.Service = (*RegisterService)(nil)

type RegisterService struct {
	repo Repository
}

func (u *RegisterService) Execute(c context.Context) (*server.Response, error) {
	data := c.Value("data")
	form, ok := data.(*model.User)
	if !ok {
		return nil, echo.ErrBadRequest
	}
	if err := u.repo.CreateUser(form); err != nil {
		return nil, err
	}
	return &server.Response{StatusCode: http.StatusCreated}, nil
}
func (u *RegisterService) RequestSchema() interface{} {
	req := new(model.User)
	return req
}
