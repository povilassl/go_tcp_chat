package hub

import (
	"net"

	"github.com/google/uuid"
)

type Client struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	DisplayName string
	Conn        net.Conn
	Send        chan *Message
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		ID:          uuid.New(),
		UserID:      uuid.Nil,
		DisplayName: "",
		Conn:        conn,
		Send:        make(chan *Message, 16),
	}
}

func (c *Client) StartWriter(h *Hub) {
	defer c.Conn.Close()

	for msg := range c.Send {
		if _, err := c.Conn.Write([]byte(msg.Format())); err != nil {
			return
		}
	}
}
