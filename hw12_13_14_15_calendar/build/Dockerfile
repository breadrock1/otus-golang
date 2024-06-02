FROM golang:1.19 as build

ARG LAUNCH_BINARY

ENV CODE_DIR /go/src
ENV BIN_DIR /opt/calendar

WORKDIR ${CODE_DIR}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

RUN make build && ln -s ${CODE_DIR}/bin/* ${BIN_DIR}

FROM alpine:3.9

LABEL ORGANIZATION="OTUS Online Education"
LABEL SERVICE="calendar"
LABEL MAINTAINERS="student@otus.ru"

COPY --from=build ${BIN_DIR}/* ${BIN_DIR}/*

ENV CONFIG_FILE /etc/calendar/config.toml
COPY ./configs/config.toml ${CONFIG_FILE}

CMD ${LAUNCH_BINARY} -config ${CONFIG_FILE}
