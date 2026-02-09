package hub

import (
	"fmt"
	"net"
	"sync/atomic"
)

var nextConnID uint64 = 0

type Client struct {
	ID   uint64
	Name string
	Conn net.Conn
	Send chan *Message
}

func NewClient(conn net.Conn) *Client {
	id := atomic.AddUint64(&nextConnID, 1)

	return &Client{
		ID:   id,
		Name: "Guest" + fmt.Sprint(id),
		Conn: conn,
		Send: make(chan *Message, 16),
	}
}

func (c *Client) StartWriter(h *Hub) {
	for msg := range c.Send {
		formattedMessage := formatMessage(msg)
		if _, err := c.Conn.Write([]byte(formattedMessage)); err != nil {
			h.Disconnect(c, "write error")
			return
		}
	}
}

func (c *Client) Rename(newName string) {
	c.Name = newName
}
