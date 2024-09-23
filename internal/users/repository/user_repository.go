package repository

import (
	"context"

	"github.com/jumayevgadam/todo_app-fiber/internal/connection"
	userModel "github.com/jumayevgadam/todo_app-fiber/internal/models/users"
	"github.com/jumayevgadam/todo_app-fiber/pkg/errlist"
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
		return -1, errlist.ParseSqlErrors(err)
	}

	return userID, nil
}
