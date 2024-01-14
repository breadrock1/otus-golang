package main

import (
	"context"
	"log"
	"net"
	"sync"
	"time"
)

func main() {
	// Place your code here,
	// P.S. Do not rush to throw context down, think think if it is useful with blocking operation?
	wg := &sync.WaitGroup{}
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second)
	wg.Add(1)
	go dealLongWithCtx(wg, ctx)
	wg.Wait()

	l, err := net.Listen("tcp", "0.0.0.0:3302")
	if err != nil {
		log.Fatalf("Cannot listen: %v", err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("Cannot accept: %v", err)
		}

		//go handleConnection(conn)
	}
}
