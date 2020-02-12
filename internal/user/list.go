package user

import (
	"context"
	"github.com/songfei1983/go-api-server/internal/server"
	"net/http"
)

func NewList(a *server.API) server.Service {
	r := NewUserRepository(a)
	return &ListService{repo: r}
}

var _ server.Service = (*ListService)(nil)

type ListService struct {
	repo Repository
}

func (u *ListService) Execute(c context.Context) (*server.Response, error) {
	data, err := u.repo.GetAllUser()
	return &server.Response{
		StatusCode: http.StatusOK,
		Data:       data,
	}, err
}
func (u *ListService) RequestSchema() interface{} {
	req := struct{}{}
	return &req
}
