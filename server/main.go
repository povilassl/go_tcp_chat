package main

import (
	"fmt"
	"net"

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

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go h.HandleConnection(conn)
	}
}
