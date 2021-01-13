package benchmark

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

// TsdbConnection represents a TimescaleDB instance, that will be the target of the
// query benchmarks.
type TsdbConnection struct {
	TsdbConnString string
}

// GetConnectionPool is resposible for opening a connection pool within the given
// TimescaleDB instance, and returning a connection pool pointer that will be used by
// the concurrent workers.
func (tc TsdbConnection) GetConnectionPool() (*pgxpool.Pool, error) {
	conn, err := pgxpool.Connect(context.Background(), tc.TsdbConnString)
	if err != nil {
		return &pgxpool.Pool{}, err
	}
	return conn, nil
}
