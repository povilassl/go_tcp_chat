package hub

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Hub struct {
	clients    map[uint64]*Client
	connect    chan *Client
	disconnect chan DisconnectEvent
	broadcast  chan Message
	send       chan Message
	execute    chan Command
}

type DisconnectEvent struct {
	Client *Client
	Reason string
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[uint64]*Client),
		connect:    make(chan *Client),
		disconnect: make(chan DisconnectEvent),
		broadcast:  make(chan Message),
		send:       make(chan Message),
		execute:    make(chan Command),
	}
}

func (h *Hub) Connect(c *Client) {
	h.connect <- c
}

func (h *Hub) Disconnect(c *Client, reason string) {
	h.disconnect <- DisconnectEvent{
		Client: c,
		Reason: reason,
	}
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
		Text: welcomeMessage,
		To:   client,
		Type: MessageSystem,
	})
}

func (h *Hub) Run() {
	for {
		select {
		case c := <-h.connect:
			h.handleConnect(c)

		case ev := <-h.disconnect:
			h.handleDisconnect(ev)

		case msg := <-h.broadcast:
			h.handleBroadcast(msg)

		case msg := <-h.send:
			h.handleSend(msg)

		case cmd := <-h.execute:
			h.handleExecute(cmd)
		}
	}
}

func (h *Hub) handleConnect(c *Client) {
	h.clients[c.ID] = c
}

func (h *Hub) handleDisconnect(ev DisconnectEvent) {

	c := ev.Client

	existingClient, ok := h.clients[ev.Client.ID]
	if !ok {
		return
	}

	fmt.Printf("Disconnecting client with ID %d. Reason: %s\r\n", c.ID, ev.Reason)

	close(existingClient.Send)
	existingClient.Conn.Close()
	delete(h.clients, c.ID)
}

func (h *Hub) handleBroadcast(msg Message) {
	for _, c := range h.clients {
		if msg.From != nil && c.ID == msg.From.ID {
			continue
		}

		select {
		case c.Send <- &msg:
		default:
			fmt.Println("Disconnecting slow client. Client ID:", c.ID)
			h.handleDisconnect(DisconnectEvent{
				Client: c,
				Reason: "slow",
			})
		}
	}
}

func (h *Hub) handleSend(msg Message) {
	if msg.To != nil {
		if c, ok := h.clients[msg.To.ID]; ok {
			select {
			case c.Send <- &msg:
			default:
				fmt.Println("Disconnecting slow client. Client ID:", c.ID)
				h.handleDisconnect(DisconnectEvent{
					Client: c,
					Reason: "slow",
				})
			}
		}
	}
}

func (h *Hub) handleExecute(cmd Command) {

	switch {
	case strings.HasPrefix(cmd.Text, "/name"):
		args := strings.SplitN(cmd.Text, " ", 2)
		if len(args) != 2 {
			h.handleSend(Message{
				Text: "Incorrect number of arguments. Usage: /name <new_name>",
				To:   cmd.From,
				Type: MessageSystem,
			})

			return
		}

		originalName := cmd.From.Name
		newName := strings.TrimSpace(args[1])

		// Validate new name
		if len(newName) == 0 || len(newName) > 14 {
			h.handleSend(Message{
				Text: "Name must be between 1 and 14 characters long",
				To:   cmd.From,
				Type: MessageSystem,
			})
			return
		}

		if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(newName) {
			h.handleSend(Message{
				Text: "Name must contain only letters and numbers",
				To:   cmd.From,
				Type: MessageSystem,
			})
			return
		}

		cmd.From.Rename(newName)

		h.handleBroadcast(Message{
			Text: fmt.Sprintf("%s is now known as %s", originalName, newName),
			Type: MessageSystem,
		})

	case strings.HasPrefix(cmd.Text, "/msg"):
		args := strings.SplitN(cmd.Text, " ", 3)
		if len(args) != 3 {
			h.handleSend(Message{
				Text: "Incorrect number of arguments. Usage: /msg <name> <message>",
				To:   cmd.From,
				Type: MessageSystem,
			})

			return
		}

		var clientName = args[1]
		var messageText = args[2]

		existingClient := h.findClientByName(clientName)
		if existingClient == nil {
			h.handleSend(Message{
				Text: "Client '" + clientName + "' is not currently online",
				To:   cmd.From,
				Type: MessageSystem,
			})
		}

		h.handleSend(Message{
			Text: messageText,
			To:   existingClient,
			From: cmd.From,
			Type: MessageDirect,
		})

	case strings.HasPrefix(cmd.Text, "/quit"):
		args := strings.SplitN(cmd.Text, " ", 2)

		exitMessage := fmt.Sprintf("%s left the server", cmd.From.Name)

		if len(args) == 2 {
			exitMessage += fmt.Sprintf(". Goodbye message: %s", args[1])
		}

		h.handleBroadcast(Message{
			Text: exitMessage,
			Type: MessageSystem,
		})

		h.handleDisconnect(DisconnectEvent{
			Client: cmd.From,
			Reason: "requested",
		})

	case strings.HasPrefix(cmd.Text, "/help"):
		h.handleSend(Message{
			Text: "Available commands: TODO",
			To:   cmd.From,
			Type: MessageSystem,
		})

	default:
		h.handleSend(Message{
			Text: "Unknown command, send '/help' for a list of available commands",
			To:   cmd.From,
			Type: MessageSystem,
		})
	}
}

func (h *Hub) findClientByName(s string) *Client {
	for _, c := range h.clients {
		if c.Name == s {
			return c
		}
	}

	return nil
}

func (h *Hub) Shutdown() {
	fmt.Println("Disconnecting all clients...")

	for _, c := range h.clients {

		c.Send <- &Message{
			Text: "Server shutting down...",
			Type: MessageSystem,
		}

		h.handleDisconnect(DisconnectEvent{
			Client: c,
			Reason: "server shutdown",
		})
	}
}
