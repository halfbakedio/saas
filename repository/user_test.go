package repository_test

import (
	"testing"

	"github.com/halfbakedio/saas/db"
	"github.com/halfbakedio/saas/ent"
	"github.com/halfbakedio/saas/ent/enttest"
	"github.com/halfbakedio/saas/repository"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

var (
	dsn = "postgres://postgres:postgres@:5432/saas_test?search_path=public&sslmode=disable"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	conn *db.Connection
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

func (s *UserRepositoryTestSuite) SetupSuite() {
	s.conn = db.GetConnection()
	s.conn.Client = enttest.Open(s.T(), "postgres", dsn)
}

func (s *UserRepositoryTestSuite) TearDownSuite() {
	s.conn.Client.Close()
}

func (s *UserRepositoryTestSuite) createTestUser(user *ent.User) (*ent.User, error) {
	repo := repository.NewUserRepository(s.conn)
	newUser, err := repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	s.T().Cleanup(func() {
		_ = repo.DeleteUser(newUser)
	})

	return newUser, nil
}

func (s *UserRepositoryTestSuite) Test_CreateUser() {
	user := &ent.User{
		Email: "foo@example.com",
	}

	newUser, err := s.createTestUser(user)

	s.NoError(err)
	s.NotNil(newUser)
	s.Equal(user.Email, newUser.Email)
}

func (s *UserRepositoryTestSuite) Test_QueryUserByName() {
	user := &ent.User{
		Email: "foo@example.com",
	}

	repo := repository.NewUserRepository(s.conn)
	_, err := s.createTestUser(user)

	s.NoError(err)

	newUser, err := repo.QueryUserByEmail(user.Email)

	s.NoError(err)
	s.NotNil(newUser)
	s.Equal(user.Email, newUser.Email)
}
