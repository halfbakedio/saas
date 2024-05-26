package service

import (
	"context"

	"github.com/halfbakedio/saas/ent"
	"github.com/halfbakedio/saas/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(client *ent.Client) *UserService {
	return &UserService{
		repo: repository.NewUserRepository(client),
	}
}

func (s *UserService) CreateUser(ctx context.Context, email string) (*ent.User, error) {
	return s.repo.CreateUser(
		ctx,
		&ent.User{
			Email: email,
		},
	)
}
