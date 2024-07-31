package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"github.org/eventmodeling/product-management/internal/infra/config"
)

var (
	db *pgxpool.Pool
)

func configureDB(ctx context.Context, host, port, user, password, database string) {

	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, database)
	db, err = config.NewPG(ctx, connectionString)

	handleError(err, "failed to create Postgresql connection")

	err = db.Ping(ctx)
	handleError(err, "failed to ping Postgresql connection")

	logrus.Info("Using Postgresql")

}
