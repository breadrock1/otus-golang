BIN_CALENDAR := "./bin/calendar"
BIN_SCHEDULER := "./bin/calendar-scheduler"
BIN_SENDER := "./bin/calendar-sender"

PROTO_BIN="~/go/bin/protoc-gen-go"
DOCKER_IMG="calendar:develop"

PG_HOST := $(if ${POSTGRES_HOST},${POSTGRES_HOST},127.0.0.1)
PG_PORT := $(if ${POSTGRES_PORT},${POSTGRES_PORT},5432)
PG_USER := $(if ${POSTGRES_USERNAME},${POSTGRES_USERNAME},postgres)
PG_PASS := $(if ${POSTGRES_PASSWORD},${POSTGRES_PASSWORD},pas)

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o $(BIN_CALENDAR) -ldflags "$(LDFLAGS)" ./cmd/calendar
	go build -v -o $(BIN_SENDER) -ldflags "$(LDFLAGS)" ./cmd/calendar_sender
	go build -v -o $(BIN_SCHEDULER) -ldflags "$(LDFLAGS)" ./cmd/calendar_scheduler

version: build
	$(BIN_CALENDAR) -V

run: build
	$(BIN_CALENDAR) -c ./configs/config.toml


build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG) \
		-f build/Dockerfile .

run-img: build-img
	docker run $(DOCKER_IMG)


migrate:
	 goose postgres "dbname=calendar sslmode=disable host=${PG_HOST} port=${PG_PORT} password=${PG_PASS} user=${PG_USER}" up

test:
	go test -race ./...


install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.50.1

lint: install-lint-deps
	golangci-lint run ./...

lint-fix: install-lint-deps
	golangci-lint run ./... --fix


install-gen-deps:
	(which protoc-gen-go > /dev/null) || go install google.golang.org/protobuf/cmd/protoc-gen-go
	(which protoc-gen-go-grpc > /dev/null) || go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	(which protoc-gen-grpc-gateway > /dev/null) || go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	(which protoc-gen-openapiv2 > /dev/null) || go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

generate: install-gen-deps
	go generate ./...

.PHONY: build run build-img run-img version test lint
