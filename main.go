package main

import (
	"fmt"
	"go-mpnj/api"
	"net/http"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/adapter/sqlite3"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

var (
	logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "main")))
)

func main() {
	var (
		// ctx        = context.Background()
		port       = 3000
		repository = initRepository()
		mux        = api.NewMux(repository)
		server     = http.Server{
			Addr:    ":3000",
			Handler: mux,
		}
	)

	fmt.Println(port)

	logger.Info("server starting: http://localhost" + server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logger.Fatal("server error", zap.Error(err))
	}
}

func initRepository() rel.Repository {
	var (
		logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "repository")))
		dsn       = "mpnj.db?_foreign_keys=1&_loc=Local"
	)

	adapter, err := sqlite3.Open(dsn)
	if err != nil {
		logger.Fatal(err.Error(), zap.Error(err))
	}

	repository := rel.New(adapter)

	return repository
}
