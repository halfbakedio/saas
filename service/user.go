package service

import (
	"github.com/halfbakedio/saas/db"
	"github.com/halfbakedio/saas/ent"
	"github.com/halfbakedio/saas/repository"
)

// UserService is a service for users.
type UserService struct {
	repo *repository.UserRepository
}

// NewUserService creates a new user service with the given client.
func NewUserService(conn *db.Connection) *UserService {
	return &UserService{
		repo: repository.NewUserRepository(conn),
	}
}

// CreateUser creates a new user in the database and returns the created user
// or an error.
func (s *UserService) CreateUser(email string) (*ent.User, error) {
	return s.repo.CreateUser(
		&ent.User{
			Email: email,
		},
	)
}

// GetUserById queries the database for a user by its ID and returns the user
// or an error.
func (s *UserService) GetUserById(id int) (*ent.User, error) {
	return s.repo.QueryUserById(id)
}

// GetUserByEmail queries the database for a user by its email and returns the
// user or an error.
func (s *UserService) GetUserByEmail(email string) (*ent.User, error) {
	return s.repo.QueryUserByEmail(email)
}
