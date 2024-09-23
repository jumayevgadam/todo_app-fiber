package connection

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // import mysql driver

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
	Execute(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
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

// NewDBConnection is
func NewDBConnection(_ context.Context, cfgs config.MySQL) (*Database, error) {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	))
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
	return d.db.ExecContext(context.Background(), query, args...)
}

// Begin is
func (d *Database) Begin(ctx context.Context, opts *sql.TxOptions) (TxOps, error) {
	tx, err := d.db.BeginTxx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("connection.Database.begin: %v", err.Error())
	}

	return &Transaction{Tx: tx}, nil
}

// Close is
func (d *Database) Close() error {
	return d.db.Close()
}
