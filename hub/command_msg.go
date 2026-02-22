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

	existingClient := h.findClientByName(clientName)
	if existingClient == nil {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: Client '%s' is not currently online", c.BaseErrorMessage(), clientName),
		)

		return
	}

	if existingClient.UserID == cmd.From.UserID {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: You cannot send a message to yourself", c.BaseErrorMessage()),
		)

		return
	}

	if _, err := h.messageService.Create(cmd.From.UserID, &existingClient.UserID, nil, messageText); err != nil {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: %s", c.BaseErrorMessage(), err.Error()),
		)
		return
	}

	msg := Message{
		Text: messageText,
		From: cmd.From,
		Type: MessageDirect,
	}

	h.sendToUserID(existingClient.UserID, msg)
}
