package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/povilassl/tcp_chat/server/connection"
	"github.com/povilassl/tcp_chat/server/hub"
)

func main() {
	h := hub.NewHub()
	go h.Run()

	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	defer ln.Close()

	fmt.Println("Listening on port 8000...")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()
		fmt.Println("Shutting down server...")
		h.Shutdown()
		ln.Close()
	}()

	for {
		conn, err := ln.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				fmt.Println("Server stopped accepting new connections.")
				return
			default:
				panic(err)
			}
		}

		go connection.Handle(h, conn)
	}
}
