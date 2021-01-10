package benchmark

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TsdbConnection struct {
	TsdbConnString string
}

func (tc TsdbConnection) GetConnectionPool() (*pgxpool.Pool, error) {
	conn, err := pgxpool.Connect(context.Background(), tc.TsdbConnString)
	if err != nil {
		return &pgxpool.Pool{}, err
	}
	return conn, nil
}
