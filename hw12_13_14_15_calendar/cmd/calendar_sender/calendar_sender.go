package main

import (
	"context"
	"encoding/json"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/cmd"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/logger"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/rabbit"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

var configFile string

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

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	_ = rabbitService.Consume(ctx, func(msg amqp.Delivery) {
		m := rabbit.Message{}
		err := json.Unmarshal(msg.Body, &m)
		if err != nil {
			log.Errorf("failed to parse bytes: %s", err)
			cancel()
			return
		}
		log.Printf("sending message %v", m)
	})
}
