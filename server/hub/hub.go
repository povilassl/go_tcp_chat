package hub

import (
	"fmt"
	"time"
)

type Hub struct {
	clients    map[uint64]*Client
	channels   map[uint64]*Channel
	commands   map[string]CommandHandler
	connect    chan *Client
	disconnect chan DisconnectEvent
	broadcast  chan Message
	send       chan Message
	execute    chan Command
}

func NewHub() *Hub {
	hub := &Hub{
		clients:    make(map[uint64]*Client),
		channels:   make(map[uint64]*Channel),
		commands:   make(map[string]CommandHandler),
		connect:    make(chan *Client),
		disconnect: make(chan DisconnectEvent),
		broadcast:  make(chan Message),
		send:       make(chan Message),
		execute:    make(chan Command),
	}

	hub.registerCommands()
	return hub
}

func (h *Hub) registerCommands() {
	h.register(&NameCommand{})
	h.register(&MsgCommand{})
	h.register(&QuitCommand{})
	h.register(&HelpCommand{})
	h.register(&CreateCommand{})
	h.register(&DeleteCommand{})
	// h.register(&JoinCommand{})
	// h.register(&LeaveCommand{})
}

func (h *Hub) register(cmd CommandHandler) {
	h.commands[cmd.Name()] = cmd
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

func (h *Hub) SendSystem(client *Client, text string) {
	h.handleSend(Message{
		Text: text,
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

func (h *Hub) sendSystem(to *Client, text string) {
	h.handleSend(Message{
		Text: text,
		To:   to,
		Type: MessageSystem,
	})
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
