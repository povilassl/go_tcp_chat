package hub

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"time"
)

func (h *Hub) HandleConnection(conn net.Conn) {

	conn.SetReadDeadline(time.Now().Add(5 * time.Minute))

	c := NewClient(conn)
	go c.StartWriter(h)
	h.Connect(c)

	fmt.Println("New client connected. Client ID:", c.ID)
	h.SendGreeting(c)

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		h.HandleMessage(scanner.Text(), c)
	}

	handleEndOfConnection(scanner.Err(), c, h)
}

func handleEndOfConnection(err error, client *Client, h *Hub) {
	if errors.Is(err, io.EOF) {
		fmt.Println("Client disconnected (EOF). Client ID:", client.ID)
	} else if ne, ok := err.(net.Error); ok && ne.Timeout() {
		fmt.Println("Read timeout, closing. Client ID:", client.ID)
	} else {
		fmt.Println("Read error, closing. ClientId:", client.ID, "Error:", err)
	}

	h.Disconnect(client)
}
