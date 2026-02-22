package hub

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type QuitCommand struct{}

func (c *QuitCommand) Name() string { return "quit" }

func (c *QuitCommand) Usage() string { return "/quit [goodbye_message]" }

func (c *QuitCommand) BaseErrorMessage() string { return "" }

func (c *QuitCommand) Execute(h *Hub, cmd Command) {

	members, err := h.channelService.GetMembersByUserID(cmd.From.UserID)
	if err == nil {
		exitMessage := ""
		if cmd.From.UserID != uuid.Nil {
			exitMessage = fmt.Sprintf("%s left the server", cmd.From.DisplayName)

			if strings.TrimSpace(cmd.Args) != "" {
				exitMessage += fmt.Sprintf(". Goodbye message: %s", cmd.Args)
			}
		}
		msg := Message{
			Text: exitMessage,
			Type: MessageSystem,
		}

		h.sendToUserIDs(*members, msg, &cmd.From.UserID)
	}

	h.handleDisconnect(DisconnectEvent{
		Client: cmd.From,
		Reason: "requested",
	})
}
