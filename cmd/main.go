package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/povilassl/tcp_chat/connection"
	"github.com/povilassl/tcp_chat/hub"
	"github.com/povilassl/tcp_chat/internal/application"
	"github.com/povilassl/tcp_chat/internal/infrastructure/db"
	"github.com/povilassl/tcp_chat/internal/infrastructure/mysql"
)

func main() {
	if err := db.RunMigrations(); err != nil {
		panic(err)
	} else {
		fmt.Println("Database migrations completed successfully.")
	}

	dbConn, err := db.NewConnection()
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	userRepo := mysql.NewUserRepository(dbConn)
	authService := application.NewAuthService(userRepo)

	h := hub.NewHub(authService)
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
