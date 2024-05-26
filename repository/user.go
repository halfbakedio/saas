package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/halfbakedio/saas/ent"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user *ent.User) (*ent.User, error)
}

type UserRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) *UserRepository {
	return &UserRepository{
		client: client,
	}
}

func (repo *UserRepository) CreateUser(ctx context.Context, user *ent.User) (*ent.User, error) {
	u, err := repo.client.User.
		Create().
		SetEmail(user.Email).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	log.Println("user was created: ", u)

	return u, nil
}
