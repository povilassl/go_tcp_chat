package hub

import (
	"fmt"
	"strings"
	"time"
)

type Hub struct {
	clients    map[uint64]*Client
	connect    chan *Client
	disconnect chan *Client
	broadcast  chan Message
	send       chan Message
	execute    chan Command
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[uint64]*Client),
		connect:    make(chan *Client),
		disconnect: make(chan *Client),
		broadcast:  make(chan Message),
		send:       make(chan Message),
		execute:    make(chan Command),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case c := <-h.connect:
			h.clients[c.ID] = c

		case c := <-h.disconnect:
			if existingClient, ok := h.clients[c.ID]; ok {
				close(existingClient.Send)
				existingClient.Conn.Close()
				delete(h.clients, c.ID)
			}

		case msg := <-h.broadcast:
			for _, c := range h.clients {
				if c.ID == msg.From.ID {
					continue
				}
				select {
				case c.Send <- &msg:
				default:
					disconnectSlowClient(c, h)
				}
			}

		case msg := <-h.send:
			if msg.To > 0 {
				if c, ok := h.clients[msg.To]; ok {
					select {
					case c.Send <- &msg:
					default:
						disconnectSlowClient(c, h)
					}
				}
			}

		case cmd := <-h.execute:
			h.ExecuteCommand(cmd)
		}
	}
}

func (h *Hub) Connect(c *Client) {
	h.connect <- c
}

func (h *Hub) Disconnect(c *Client) {
	h.disconnect <- c
}

func (h *Hub) Broadcast(msg Message) {
	h.broadcast <- msg
}

func (h *Hub) SendDirect(msg Message) {
	h.send <- msg
}

func (h *Hub) Execute(cmd Command) {
	h.execute <- cmd
}

func (h *Hub) SendGreeting(client *Client) {
	welcomeMessage := "Welcome to the server!\r\n\r\n"
	welcomeMessage += "-----------------------------\r\n"
	welcomeMessage += "Your client ID is: " + fmt.Sprint(client.ID) + "\r\n"
	welcomeMessage += "Your name is: " + client.Name + "\r\n"
	welcomeMessage += "Time of connection: " + time.Now().Format(time.RFC1123) + "\r\n"
	welcomeMessage += "-----------------------------\r\n"

	h.SendDirect(Message{
		Text:   welcomeMessage,
		To:     client.ID,
		System: true,
	})
}

func (h *Hub) HandleMessage(s string, c *Client) {
	if strings.HasPrefix(s, "/") {
		h.Execute(Command{
			Text: s,
			From: c})
	} else {
		h.Broadcast(Message{
			Text:   s,
			From:   c,
			System: false,
		})
	}
}

func disconnectSlowClient(c *Client, h *Hub) {
	fmt.Println("Disconnecting slow client. Client ID:", c.ID)
	h.disconnect <- c
}
