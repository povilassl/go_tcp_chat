package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/jmoiron/sqlx"
	"github.com/povilassl/tcp_chat/connection"
	"github.com/povilassl/tcp_chat/hub"
	"github.com/povilassl/tcp_chat/internal/application"
	"github.com/povilassl/tcp_chat/internal/infrastructure/db"
	"github.com/povilassl/tcp_chat/internal/infrastructure/mysql"
)

func main() {
	if err := db.RunMigrations(); err != nil {
		fatal("Migration error", err)
	}

	fmt.Println("Database migrations completed successfully.")

	dbConn, err := db.NewConnection()
	if err != nil {
		fatal("Database connection error", err)
	}

	defer dbConn.Close()

<<<<<<< HEAD
	h := bootstrapHub(dbConn)
=======
	userRepo := mysql.NewUserRepository(dbConn)
	channelRepo := mysql.NewChannelRepository(dbConn)
	messageRepo := mysql.NewMessageRepository(dbConn)

	authService := application.NewAuthService(userRepo)
	channelservice := application.NewChannelService(channelRepo, messageRepo)
	messageService := application.NewMessageService(messageRepo, channelRepo)
	userService := application.NewUserService(userRepo)

	h := hub.NewHub(authService, channelservice, messageService, userService)
>>>>>>> 956b4c2fd92979f87329a0c76968fbd5b5aa4fa4
	go h.Run()

	port := getServerPort()
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fatal("Listen error", err)
	}

	defer ln.Close()

	fmt.Printf("Listening on port %s...\n", port)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()
		fmt.Println("Shutdown signal received, closing listener...")
		ln.Close()
	}()

	for {
		conn, err := ln.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				fmt.Println("Shutting down server...")
				h.Shutdown()
				fmt.Println("Server stopped.")
				return
			default:
				fmt.Fprintf(os.Stderr, "Accept error: %v\n", err)
			}
		}

		go connection.Handle(h, conn)
	}
}

func bootstrapHub(dbConn *sqlx.DB) *hub.Hub {
	userRepo := mysql.NewUserRepository(dbConn)
	channelRepo := mysql.NewChannelRepository(dbConn)
	messageRepo := mysql.NewMessageRepository(dbConn)

	authService := application.NewAuthService(userRepo)
	channelService := application.NewChannelService(channelRepo, messageRepo)
	messageService := application.NewMessageService(messageRepo, channelRepo)
	userService := application.NewUserService(userRepo)

	return hub.NewHub(authService, channelService, messageService, userService)
}

func getServerPort() string {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8000"
	}
	return port
}

func fatal(msg string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %v\n", msg, err)
	os.Exit(1)
}
