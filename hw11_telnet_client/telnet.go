package main

import (
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type implClient struct {
	address    string
	timeout    time.Duration
	connection net.Conn
	isClosed   bool
	stdIn      io.Reader
	stdOut     io.Writer
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &implClient{
		address: address,
		timeout: timeout,
		stdIn:   in,
		stdOut:  out,
	}
}

func (i implClient) Connect() error {
	//TODO implement me
	panic("implement me")
}

func (i implClient) Close() error {
	//TODO implement me
	panic("implement me")
}

func (i implClient) Send() error {
	//TODO implement me
	panic("implement me")
}

func (i implClient) Receive() error {
	//TODO implement me
	panic("implement me")
}
