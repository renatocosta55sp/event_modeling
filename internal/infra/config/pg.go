package config

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5"
)

var (
	pgInstance *pgx.Conn
	pgOnce     sync.Once
)

func NewPG(ctx context.Context, connString string) (*pgx.Conn, error) {
	var err error

	pgOnce.Do(func() {
		var db *pgx.Conn
		db, err = pgx.Connect(ctx, connString)
		if err == nil {
			pgInstance = db
		}
	})

	return pgInstance, err
}
