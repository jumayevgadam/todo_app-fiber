package connection

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/jumayevgadam/todo_app-fiber/internal/config"
)

// DBops is
var _ DB = (*Database)(nil)

// DB is
type DB interface {
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row
	Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

// DBops is
type DBops interface {
	DB
	Begin(ctx context.Context, opts *sql.TxOptions) (TxOps, error)
	Close() error
}

// Database struct is
type Database struct {
	db *sqlx.DB
}

// GenerateDsn is
func GenerateDsn(cfgs config.MySQL) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfgs.User,
		cfgs.Password,
		cfgs.Host,
		cfgs.Port,
		cfgs.DBName,
	)
}

// NewDBConnection is
func NewDBConnection(_ context.Context, cfgs config.MySQL) (*Database, error) {
	db, err := sqlx.Connect("mysql", GenerateDsn(cfgs))
	if err != nil {
		return nil, fmt.Errorf("db.Connect: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("db.Ping: %w", err)
	}

	return &Database{db: db}, nil
}

// Get is
func (d *Database) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return d.db.GetContext(ctx, dest, query, args...)
}

// QueryRow is
func (d *Database) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return d.db.QueryRowContext(ctx, query, args...)
}

// Query is
func (d *Database) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return d.db.QueryContext(ctx, query, args...)
}

// Select is
func (d *Database) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return d.db.SelectContext(ctx, dest, query, args...)
}

// Execute is
func (d *Database) Execute(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return d.db.ExecContext(ctx, query, args...)
}

// Begin is
func (d *Database) Begin(ctx context.Context, opts *sql.TxOptions) (TxOps, error) {
	t, err := d.db.BeginTxx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &Transaction{Tx: t}, nil
}

// Close is
func (d *Database) Close() error {
	return d.db.Close()
}
