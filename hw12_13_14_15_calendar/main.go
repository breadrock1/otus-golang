package main

import (
	"context"
	"fmt"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/cmd"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/app"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/server/http"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/storage"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/storage/memcache"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/storage/sqlstorage"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config := cmd.Execute()

	sLog, err := logger.New(&config.Logger)
	if err != nil {
		log.Fatalln(err)
	}

	var storageService storage.Storage
	if !config.Database.EnableInMemory {
		storageService = sqlstorage.New()
	} else {
		storageService = memcache.New()
	}

	address := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.HostPort)
	calendar := app.New(storageService, sLog)
	server := internalhttp.NewServer(address, calendar, sLog)
	signals := []os.Signal{syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP}
	ctx, cancel := signal.NotifyContext(context.Background(), signals...)
	defer cancel()

	if err = server.Start(ctx); err != nil {
		sLog.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1)
	}
}
