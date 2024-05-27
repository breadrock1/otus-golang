package main

import (
	"context"
	"encoding/json"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/cmd"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/logger"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/rabbit"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/storage"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/storage/event"
	log "github.com/sirupsen/logrus"
)

var configFile string

const (
	removeTimeout = time.Minute * 5
	checkTimout   = time.Minute
)

func newMessage(ev event.Event) rabbit.Message {
	return rabbit.Message{
		ID:     ev.ID,
		Name:   ev.Title,
		Time:   ev.Start,
		UserID: ev.UserID,
	}
}

func init() {
	flag.StringVar(&configFile, "config", "./configs/config.toml", "Path to configuration file")
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

func main() {
	config := cmd.Execute()

	_, err := logger.New(&config.Logger)
	if err != nil {
		log.Fatalln(err)
	}

	rabbitService := rabbit.New(config.Rabbit)
	if err = rabbitService.Connect(); err != nil {
		log.Fatalln(err)
	}
	defer rabbitService.Close()

	storageService := storage.New(&config.Database)

	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		if err := storageService.Close(ctx); err != nil {
			log.Println(err)
		}
	}()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	startTime := time.Now().Add(-time.Minute)
	endTime := time.Now()
	checkTicker := time.NewTicker(checkTimout)
	removeTicker := time.NewTicker(removeTimeout)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			log.Debugf("get events: %s - %s", startTime, endTime)
			events, err := storageService.GetEventsByNotifier(ctx, startTime, endTime)
			if err != nil {
				log.Errorf("failed to get events: %s", err)
				continue
			}

			for _, ev := range events {
				log.Debugf("send event: %v", ev)
				m := newMessage(ev)
				data, _ := json.Marshal(m)
				_ = rabbitService.Publish(data)
			}

			select {
			case <-ctx.Done():
				return
			case <-checkTicker.C:
				log.Debug("ticker")
				startTime = endTime
				endTime = time.Now()
			case <-removeTicker.C:
				storageService.RemoveAfter(ctx, time.Now().Add(-1*(time.Hour*24*365)))
			}
		}
	}
}
