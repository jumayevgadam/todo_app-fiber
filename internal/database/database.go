package database

import (
	"context"

	"github.com/jumayevgadam/todo_app-fiber/internal/users"
)

// Transaction is
type Transaction func(db DataStore) error

// DataStore is
type DataStore interface {
	WithTransaction(ctx context.Context, transaction Transaction) error
	UsersRepo() users.Repository
}
