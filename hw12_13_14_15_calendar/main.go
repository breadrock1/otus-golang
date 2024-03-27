package main

import (
	"context"
	"flag"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/cmd"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/app"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/server/http"
	memcache "github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/storage/memcache"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config := cmd.Execute()

	flag.Parse()
	if flag.Arg(0) == "version" {
		cmd.PrintVersion()
		return
	}

	sLog, err := logger.New(config)
	if err != nil {
		log.Fatalln(err)
	}

	storage := memcache.Storage{}
	calendar := app.New(sLog, storage)

	server := internalhttp.NewServer(sLog, calendar)

	signals := []os.Signal{syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP}
	ctx, cancel := signal.NotifyContext(context.Background(), signals...)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			sLog.Error("failed to stop http server: " + err.Error())
		}
	}()

	sLog.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		sLog.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1)
	}
}
