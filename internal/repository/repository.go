package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/songfei1983/go-api-server/ent"
)

type Repository struct {
	client *ent.Client
}

func New(client *ent.Client) *Repository {
	return &Repository{client}
}

func (r *Repository) CreateUser(ctx context.Context, name string) (*ent.Adunit, error) {
	u, err := r.client.Adunit.
		Create().
		SetName(name).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func (r *Repository) QueryUser(ctx context.Context) ([]*ent.Adunit, error) {
	u, err := r.client.Adunit.
		Query().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %v", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}
