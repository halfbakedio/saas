package repository

import (
	"fmt"
	"log"

	"github.com/halfbakedio/saas/db"
	"github.com/halfbakedio/saas/ent"
	"github.com/halfbakedio/saas/ent/user"
)

type IUserRepository interface {
	CreateUser(user *ent.User) (*ent.User, error)
	QueryUserById(id int) (*ent.User, error)
	QueryUserByEmail(email string) (*ent.User, error)
}

// UserRepository is a repository for users.
type UserRepository struct {
	conn *db.Connection
}

// NewUserRepository creates a new user repository with the given client.
func NewUserRepository(conn *db.Connection) *UserRepository {
	return &UserRepository{
		conn: conn,
	}
}

// CreateUser creates a new user in the database and returns the created user
// or an error.
func (repo *UserRepository) CreateUser(user *ent.User) (*ent.User, error) {
	u, err := repo.conn.Client.User.
		Create().
		SetEmail(user.Email).
		Save(repo.conn.Ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	log.Println("user was created: ", u)

	return u, nil
}

// QueryUserById queries the database for a user by its ID and returns the user
// or an error.
func (repo *UserRepository) QueryUserById(id int) (*ent.User, error) {
	u, err := repo.conn.Client.User.
		Query().
		Where(user.ID(id)).
		Only(repo.conn.Ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	log.Println("user returned: ", u)

	return u, nil
}

// QueryUserByEmail queries the database for a user by its email and returns
// the user or an error.
func (repo *UserRepository) QueryUserByEmail(email string) (*ent.User, error) {
	u, err := repo.conn.Client.User.
		Query().
		Where(user.Email(email)).
		Only(repo.conn.Ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	log.Println("user returned: ", u)

	return u, nil
}

// DeleteUser deletes a user from the database.
func (repo *UserRepository) DeleteUser(user *ent.User) error {
	err := repo.conn.Client.User.
		DeleteOne(user).
		Exec(repo.conn.Ctx)

	if err != nil {
		return fmt.Errorf("failed deleting user: %w", err)
	}

	log.Println("user was deleted: ", user)

	return nil
}
