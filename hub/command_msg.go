package hub

import "strings"

type MsgCommand struct{}

func (c *MsgCommand) Name() string { return "msg" }

func (c *MsgCommand) Usage() string { return "/msg <name> <message>" }

func (c *MsgCommand) Execute(h *Hub, cmd Command) {
	args := strings.SplitN(cmd.Args, " ", 2)
	if len(args) != 2 {
		h.sendSystemToClient(
			cmd.From,
			"Incorrect number of arguments. Usage: /msg <name> <message>",
		)

		return
	}

	var clientName = args[0]
	var messageText = args[1]

	existingClient := h.findClientByName(clientName)
	if existingClient == nil {
		h.sendSystemToClient(
			cmd.From,
			"Client '"+clientName+"' is not currently online",
		)

		return
	}

	h.handleSend(Message{
		Text: messageText,
		To:   existingClient,
		From: cmd.From,
		Type: MessageDirect,
	})
}
