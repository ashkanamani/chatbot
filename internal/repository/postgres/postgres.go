package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func NewPostgresConnection(address string) (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), address)
}
