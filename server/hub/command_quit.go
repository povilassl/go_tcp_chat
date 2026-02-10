package hub

import (
	"fmt"
	"strings"
)

type QuitCommand struct{}

func (c *QuitCommand) Name() string { return "quit" }

func (c *QuitCommand) Execute(h *Hub, cmd Command) {
	exitMessage := fmt.Sprintf("%s left the server", cmd.From.Name)

	if strings.TrimSpace(cmd.Args) != "" {
		exitMessage += fmt.Sprintf(". Goodbye message: %s", cmd.Args)
	}

	h.handleBroadcast(Message{
		Text: exitMessage,
		Type: MessageSystem,
	})

	h.handleDisconnect(DisconnectEvent{
		Client: cmd.From,
		Reason: "requested",
	})
}
