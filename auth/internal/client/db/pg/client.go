package pg

import (
	"context"

	"github.com/bifidokk/awesome-chat/auth/internal/client/db"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type pgClient struct {
	dbConnection db.DB
}

// New establishes a new database connection and returns a db.Client.
func New(ctx context.Context, dsn string) (db.Client, error) {
	connection, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, errors.Errorf("failed to connect to db: %v", err)
	}

	return &pgClient{
		dbConnection: NewDB(connection),
	}, nil
}

func (c *pgClient) DB() db.DB {
	return c.dbConnection
}

func (c *pgClient) Close() error {
	if c.dbConnection != nil {
		c.dbConnection.Close()
	}

	return nil
}
