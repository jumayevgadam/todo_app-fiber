package repository

import (
	"context"
	"fmt"

	"github.com/jumayevgadam/todo_app-fiber/internal/connection"
	userModel "github.com/jumayevgadam/todo_app-fiber/internal/models/users"
)

// UserRepository struct is
type UserRepository struct {
	mysqlDB connection.DB
}

// NewUserRepository is
func NewUserRepository(mysqlDB connection.DB) *UserRepository {
	return &UserRepository{mysqlDB: mysqlDB}
}

// SignUp is
func (ur *UserRepository) SignUp(ctx context.Context, resDAO *userModel.SignUpRes) (int64, error) {
	var userID int64

	err := ur.mysqlDB.QueryRow(
		ctx,
		SignUPQuery,
		resDAO.Username,
		resDAO.Email,
		resDAO.Password,
	).Scan(&userID)
	if err != nil {
		return -1, fmt.Errorf("error in repo: %w", err)
	}

	return userID, nil
}
