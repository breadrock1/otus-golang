package main

import (
	"context"
	"fmt"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/cmd"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/app"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/logger"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/server/grpcserv"
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

	host := config.Server.Host
	calendar := app.New(storageService, sLog)

	ctx, cancel := context.WithCancel(context.Background())
	go awaitSystemSignals(cancel)

	httpAddress := fmt.Sprintf("%s:%d", host, config.Server.HostPort)
	httpServer := internalhttp.NewServer(httpAddress, calendar, sLog)

	go func() {
		err := httpServer.Start(ctx)
		if err != nil {
			log.Println(err)
			cancel()
		}
	}()

	grpcAddress := fmt.Sprintf("%s:%d", host, config.Server.GrpcPort)
	grpcServer := grpcserv.NewServer(grpcAddress, calendar, sLog)

	go func() {
		err := grpcServer.Start()
		if err != nil {
			log.Println(err)
			cancel()
		}
	}()

	<-ctx.Done()
	cancel()
	shutdownServices(ctx, httpServer, grpcServer)
}

func awaitSystemSignals(cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	cancel()
}

func shutdownServices(ctx context.Context, httpServ *internalhttp.Server, grpcServ *grpcserv.Server) {
	if err := httpServ.Stop(ctx); err != nil {
		log.Println(err)
	}

	if err := grpcServ.Stop(ctx); err != nil {
		log.Println(err)
	}
}
