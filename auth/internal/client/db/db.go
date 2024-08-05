package db

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// Handler defines a function type that takes a context and returns an error.
type Handler func(ctx context.Context) error

// Client represents a database client interface.
type Client interface {
	DB() DB
	Close() error
}

// TxManager defines an interface for managing transactions with read-committed isolation level.
type TxManager interface {
	ReadCommitted(ctx context.Context, f Handler) error
}

// Query represents a SQL query with a name and raw query string.
type Query struct {
	Name     string
	QueryRaw string
}

// Transactor defines an interface for beginning transactions.
type Transactor interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

// SQLExecer combines NamedExecer and QueryExecer interfaces.
type SQLExecer interface {
	NamedExecer
	QueryExecer
}

// NamedExecer defines methods for executing named SQL queries and scanning results.
type NamedExecer interface {
	ScanOneContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	ScanAllContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

// QueryExecer defines methods for executing raw SQL queries.
type QueryExecer interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

// Pinger defines an interface for checking the connection to the database.
type Pinger interface {
	Ping(ctx context.Context) error
}

// DB combines SQLExecer, Transactor, and Pinger interfaces, and defines a method for closing the database connection.
type DB interface {
	SQLExecer
	Transactor
	Pinger
	Close()
}
