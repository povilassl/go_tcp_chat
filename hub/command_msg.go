package hub

import (
	"fmt"
)

type MsgCommand struct{}

func (c *MsgCommand) Name() string { return "msg" }

func (c *MsgCommand) Usage() string { return "/msg <name> <message>" }

func (c *MsgCommand) BaseErrorMessage() string { return "Error sending message" }

func (c *MsgCommand) Execute(h *Hub, cmd Command) {
	if !h.RequireAuth(cmd, c.BaseErrorMessage()) {
		return
	}

	args, ok := h.GetArgs(cmd, 2, c.Usage(), c.BaseErrorMessage())
	if !ok {
		return
	}

	var clientName = args[0]
	var messageText = args[1]

	//TODO continue here
	existingClient := h.findClientByName(clientName)
	if existingClient == nil {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: Client '%s' is not currently online", c.BaseErrorMessage(), clientName),
		)

		return
	}

	//TODO: save message to history

	h.handleSend(Message{
		Text: messageText,
		To:   existingClient,
		From: cmd.From,
		Type: MessageDirect,
	})
}
