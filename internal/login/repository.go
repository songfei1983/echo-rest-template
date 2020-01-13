package login

import (
	"github.com/songfei1983/go-api-server/cmd/api/app"
	"github.com/songfei1983/go-api-server/internal/persistence"
)

type Repository interface {

}

type loginPersistence struct {
	persistence.Persistence
}

func NewLoginPersistence(api app.APP) Repository{
	return &loginPersistence{persistence.Persistence{DB: api.DB}}
}
