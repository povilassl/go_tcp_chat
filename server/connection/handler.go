package connection

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"time"

	"github.com/povilassl/tcp_chat/server/hub"
)

func Handle(h *hub.Hub, conn net.Conn) {

	conn.SetReadDeadline(time.Now().Add(5 * time.Minute))

	c := hub.NewClient(conn)
	go c.StartWriter(h)
	h.Connect(c)

	fmt.Println("New client connected. Client ID:", c.ID)
	h.SendGreeting(c)

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		handleInput(h, c, scanner.Text())
	}

	handleEndOfConnection(c, h, scanner.Err())
}

func handleEndOfConnection(client *hub.Client, h *hub.Hub, err error) {
	var reason string

	if errors.Is(err, io.EOF) {
		fmt.Println("Client disconnected (EOF). Client ID:", client.ID)
		reason = "disconnected"
	} else if ne, ok := err.(net.Error); ok && ne.Timeout() {
		fmt.Println("Read timeout, closing. Client ID:", client.ID)
		reason = "timeout"
	} else {
		fmt.Println("Read error, closing. ClientId:", client.ID, "Error:", err)
		reason = "read error"
	}

	h.Disconnect(client, reason)
}

// TODO possibility to leave whitespace in messages
func handleInput(h *hub.Hub, c *hub.Client, raw string) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return
	}

	if strings.HasPrefix(raw, "/") {
		h.Execute(hub.Command{
			Text: raw,
			From: c})

		return
	}

	h.Broadcast(hub.Message{
		Text: raw,
		From: c,
		Type: hub.MessagePublic,
	})
}
