package hub

import (
	"fmt"
	"strings"
)

type MsgCommand struct{}

func (c *MsgCommand) Name() string { return "msg" }

func (c *MsgCommand) Usage() string { return "/msg <name> <message>" }

func (c *MsgCommand) BaseErrorMessage() string { return "Error sending message" }

func (c *MsgCommand) Execute(h *Hub, cmd Command) {
	args := strings.SplitN(cmd.Args, " ", 2)
	if len(args) != 2 {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: Incorrect number of arguments. Usage: %s", c.BaseErrorMessage(), c.Usage()),
		)

		return
	}

	var clientName = args[0]
	var messageText = args[1]

	existingClient := h.findClientByName(clientName)
	if existingClient == nil {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: Client '%s' is not currently online", c.BaseErrorMessage(), clientName),
		)

		return
	}

	//save message to history

	h.handleSend(Message{
		Text: messageText,
		To:   existingClient,
		From: cmd.From,
		Type: MessageDirect,
	})
}
