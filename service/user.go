package service

import (
	"errors"

	"github.com/halfbakedio/saas/db"
	"github.com/halfbakedio/saas/ent"
	"github.com/halfbakedio/saas/repository"
	"github.com/halfbakedio/saas/util"

	log "github.com/sirupsen/logrus"
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
func (s *UserService) CreateUser(email, password string) (*ent.User, error) {
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return nil, err
	}

	return s.repo.CreateUser(
		&ent.User{
			Email:    email,
			Password: hashedPassword,
		},
	)
}

// FindUserById queries the database for a user by its ID and returns the user
// or an error.
func (s *UserService) FindUserByID(id int) (*ent.User, error) {
	return s.repo.QueryUserByID(id)
}

// FindUserByEmail queries the database for a user by its email and returns the
// user or an error.
func (s *UserService) FindUserByEmail(email string) (*ent.User, error) {
	return s.repo.QueryUserByEmail(email)
}

// FindOrCreateUser queries the database for a user by its email and returns the
// user if it exists. If the user does not exist, it creates a new user and
// returns the created user or an error.
func (s *UserService) FindOrCreateUser(
	email string,
	attributes map[string]interface{}, // TODO: create Attributes type
) (*ent.User, error) {
	user, err := s.FindUserByEmail(email)
	if err != nil {
		log.Debugf("could not find user with email %s, creating a new one", email)
	}

	if user != nil {
		return user, nil
	}

	password, ok := attributes["password"].(string)
	if !ok {
		return nil, errors.New("password is required")
	}

	return s.CreateUser(email, password)
}

// DeleteUserByEmail deletes a user from the database by its ID and returns the user
// or an error.
func (s *UserService) DeleteUserByEmail(email string) error {
	log.Debugf("Deleting user with email %s", email)
	user, err := s.FindUserByEmail(email)
	if err != nil {
		return err
	}
	return s.repo.DeleteUser(user)
}
