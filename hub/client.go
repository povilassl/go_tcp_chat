package hub

import (
	"fmt"
	"net"
	"sync/atomic"

	"github.com/povilassl/tcp_chat/internal/domain/entity"
)

var nextConnectionID uint64 = 0

type Client struct {
	ID   uint64
	Name string
	Conn net.Conn
	User *entity.User
	Send chan *Message
}

func NewClient(conn net.Conn) *Client {
	id := atomic.AddUint64(&nextConnectionID, 1)

	return &Client{
		ID:   id,
		Name: "Guest" + fmt.Sprint(id),
		Conn: conn,
		User: nil,
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

func formatMessage(m *Message) string {
	var sender string

	switch m.Type {
	case MessageSystem:
		sender = "* System"

	case MessageDirect:
		sender = fmt.Sprintf("DM from %s", m.From.Name)

	case MessageChannel:
		sender = fmt.Sprintf("%s in Channel #%s", m.From.Name, m.Channel.Name)

	default:
		sender = "Unknown"
	}

	return fmt.Sprintf("[%s] %s\r\n", sender, m.Text)
}

// TODO: possibility to add persistent nicknames
func (c *Client) Login(user *entity.User) {
	c.User = user
}

func (c *Client) Rename(newName string) {
	c.Name = newName
}
