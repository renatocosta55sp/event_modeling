package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.org/napp/product-management/internal/infra/config"
	"github.org/napp/product-management/internal/infra/ports/http"
)

var (
	err error
)

func handleError(err error, msg string) {
	if err != nil {
		logrus.WithError(err).Fatal(msg)
	}
}

func main() {

	ctx := context.Background()

	ctx, stopFn := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stopFn()

	logrus.Info(ctx, config.Config.GetString("APP_NAME")+" - v.: "+config.Config.GetString("APP_VERSION"))

	configureDB(ctx,
		config.Config.GetString("DB_HOST"),
		config.Config.GetString("DB_PORT"),
		config.Config.GetString("DB_USER"),
		config.Config.GetString("DB_PASS"),
		config.Config.GetString("DB_NAME"),
	)

	defer db.Close(ctx)

	server := gin.Default()
	h := http.HttpServer{Db: db}
	http.InitRoutes(&server.RouterGroup, h)

	serverErr := make(chan error, 1)
	go func() {
		if err := server.Run(":8888"); err != nil {
			serverErr <- err
		}
	}()

	//Graceful shutdown
	select {
	case err = <-serverErr:
		handleError(err, "error on server execution")
	case <-ctx.Done():
		handleError(ctx.Err(), "ctx error")
		handleError(err, "failed to stop server gracefully")
	}

}
