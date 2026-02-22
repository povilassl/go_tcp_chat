package hub

import (
	"github.com/google/uuid"
	"github.com/povilassl/tcp_chat/internal/application/interfaces"
)

type Hub struct {
	clientsByUserID map[uuid.UUID]*Client
	pendingClients  map[uuid.UUID]*Client
	commands        map[string]CommandHandler
	connect         chan *Client
	disconnect      chan DisconnectEvent
	send            chan Message
	execute         chan Command
	authService     interfaces.AuthService
	channelService  interfaces.ChannelService
	messageService  interfaces.MessageService
	userService     interfaces.UserService
}

func NewHub(
	authService interfaces.AuthService,
	channelService interfaces.ChannelService,
	messageService interfaces.MessageService,
	userService interfaces.UserService) *Hub {
	hub := &Hub{
		clientsByUserID: make(map[uuid.UUID]*Client),
		pendingClients:  make(map[uuid.UUID]*Client),
		commands:        make(map[string]CommandHandler),
		connect:         make(chan *Client),
		disconnect:      make(chan DisconnectEvent),
		send:            make(chan Message),
		execute:         make(chan Command),
		authService:     authService,
		channelService:  channelService,
		messageService:  messageService,
		userService:     userService,
	}

	hub.registerCommands()
	return hub
}

func (h *Hub) registerCommands() {
	commands := []CommandHandler{
		&NameCommand{},
		&MsgCommand{},
		&QuitCommand{},
		&HelpCommand{},
		&CreateCommand{},
		&DeleteCommand{},
		&JoinCommand{},
		&LeaveCommand{},
		&ChannelCommand{},
		&GetCommand{},
		&RegisterCommand{},
		&LoginCommand{},
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

func (h *Hub) handleExecute(cmd Command) {
	handler, ok := h.commands[cmd.Name]
	if !ok {
		h.sendSystemToClient(
			cmd.From,
			"Unknown command. Send '/help' for a list of available commands.",
		)
		return
	}

	handler.Execute(h, cmd)
}
