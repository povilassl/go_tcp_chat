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

func handleInput(h *hub.Hub, c *hub.Client, raw string) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return
	}

	cmd := ParseCommand(raw, c)
	h.Execute(cmd)
}

func ParseCommand(line string, c *hub.Client) hub.Command {
	line = strings.TrimSpace(line)
	line = strings.TrimPrefix(line, "/")

	parts := strings.SplitN(line, " ", 2)

	cmd := hub.Command{
		Name: parts[0],
		From: c,
	}

	if len(parts) == 2 {
		cmd.Args = parts[1]
	}

	return cmd
}
