package hub

import (
	"fmt"
	"time"

	"github.com/povilassl/tcp_chat/internal/application"
)

type Hub struct {
	clients        map[uint64]*Client
	channels       map[uint64]*Channel
	commands       map[string]CommandHandler
	connect        chan *Client
	disconnect     chan DisconnectEvent
	send           chan Message
	execute        chan Command
	authService    *application.AuthService
	channelService *application.ChannelService
}

func NewHub(
	authService *application.AuthService,
	channelService *application.ChannelService) *Hub {
	hub := &Hub{
		clients:        make(map[uint64]*Client),
		channels:       make(map[uint64]*Channel),
		commands:       make(map[string]CommandHandler),
		connect:        make(chan *Client),
		disconnect:     make(chan DisconnectEvent),
		send:           make(chan Message),
		execute:        make(chan Command),
		authService:    authService,
		channelService: channelService,
	}

	hub.registerCommands()
	return hub
}

func (h *Hub) registerCommands() {
	base := &BaseCommand{}
	commands := []CommandHandler{
		&NameCommand{BaseCommand: *base},
		&MsgCommand{BaseCommand: *base},
		&QuitCommand{BaseCommand: *base},
		&HelpCommand{BaseCommand: *base},
		&CreateCommand{BaseCommand: *base},
		&DeleteCommand{BaseCommand: *base},
		&JoinCommand{BaseCommand: *base},
		&LeaveCommand{BaseCommand: *base},
		&ChannelCommand{BaseCommand: *base},
		&GetCommand{BaseCommand: *base},
		&RegisterCommand{BaseCommand: *base},
	}

	for _, cmd := range commands {
		h.commands[cmd.Name()] = cmd
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

	h.handleSend(Message{
		Text: welcomeMessage,
		To:   client,
		Type: MessageSystem,
	})
}

func (h *Hub) sendSystemGlobalBroadcast(text string) {
	h.handleBroadcast(Message{
		Text: text,
		Type: MessageSystem,
	})
}

func (h *Hub) sendSystemToClient(client *Client, text string) {
	h.handleSend(Message{
		Text: text,
		To:   client,
		Type: MessageSystem,
	})
}

func (h *Hub) sendSystemToChannel(channel *Channel, text string) {
	h.handleSend(Message{
		Text:    text,
		Channel: channel,
		Type:    MessageSystem,
	})
}

func (h *Hub) Run() {
	for {
		select {
		case c := <-h.connect:
			h.handleConnect(c)

		case ev := <-h.disconnect:
			h.handleDisconnect(ev)

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

func (h *Hub) disconnectSlowClient(c *Client) {
	fmt.Println("Disconnecting slow client. Client ID:", c.ID)
	h.handleDisconnect(DisconnectEvent{
		Client: c,
		Reason: "slow",
	})
}

func (h *Hub) handleBroadcast(msg Message) {
	for _, c := range h.clients {
		if msg.From != nil && c.ID == msg.From.ID {
			continue
		}

		select {
		case c.Send <- &msg:
		default:
			h.disconnectSlowClient(c)
		}
	}
}

func (h *Hub) handleSend(msg Message) {
	recipients := msg.getRecipients()
	for _, c := range recipients {
		select {
		case c.Send <- &msg:
		default:
			h.disconnectSlowClient(c)
		}
	}
}

func (h *Hub) handleExecute(cmd Command) {
	handler, ok := h.commands[cmd.Name]
	if !ok {
		h.handleSend(Message{
			Text: "Unknown command. Send '/help' for a list of available commands.",
			To:   cmd.From,
			Type: MessageSystem,
		})

		return
	}

	handler.Execute(h, cmd)
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
