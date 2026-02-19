package hub

import (
	"fmt"
	"strings"
)

type QuitCommand struct{}

func (c *QuitCommand) Name() string { return "quit" }

func (c *QuitCommand) Usage() string { return "/quit [goodbye_message]" }

func (c *QuitCommand) BaseErrorMessage() string { return "" }

func (c *QuitCommand) Execute(h *Hub, cmd Command) {

	if cmd.From.User != nil {
		exitMessage := fmt.Sprintf("%s left the server", cmd.From.User.Nickname)

		if strings.TrimSpace(cmd.Args) != "" {
			exitMessage += fmt.Sprintf(". Goodbye message: %s", cmd.Args)
		}
	}

	//TODO send system to channels where client was connected
	// h.handleBroadcast(Message{
	// 	Text: exitMessage,
	// 	Type: MessageSystem,
	// })

	h.handleDisconnect(DisconnectEvent{
		Client: cmd.From,
		Reason: "requested",
	})
}
