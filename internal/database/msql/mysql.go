package msql

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/jumayevgadam/todo_app-fiber/internal/connection"
	"github.com/jumayevgadam/todo_app-fiber/internal/database"
	"github.com/jumayevgadam/todo_app-fiber/internal/users"
	userRepository "github.com/jumayevgadam/todo_app-fiber/internal/users/repository"
)

var _ database.DataStore = (*DataStore)(nil)

// DataStore is
type DataStore struct {
	db       connection.DB
	user     users.Repository
	userInit sync.Once
}

// NewDataStore is
func NewDataStore(db connection.DBops) database.DataStore {
	return &DataStore{
		db: db,
	}
}

// UsersRepo is
func (d *DataStore) UsersRepo() users.Repository {
	d.userInit.Do(func() {
		d.user = userRepository.NewUserRepository(d.db)
	})

	return d.user
}

func (d *DataStore) WithTransaction(ctx context.Context, transactionFn database.Transaction) (err error) {
	db, ok := d.db.(connection.DBops)
	if !ok {
		return fmt.Errorf("got error start of transaction")
	}

	tx, err := db.Begin(ctx, nil)
	if err != nil {
		return fmt.Errorf("db.Begin: %w", err)
	}

	defer func() {
		if err != nil {
			if err = tx.Rollback(); err != nil {
				log.Printf("[mysql][WithTransaction] failed to rollback transaction: %v", err)
			}
		}
	}()

	transactionalDB := &DataStore{db: tx}
	if err := transactionFn(transactionalDB); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
