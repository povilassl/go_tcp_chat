package hub

type ChannelCommand struct {
	BaseCommand
}

func (c *ChannelCommand) Name() string { return "channel" }

func (c *ChannelCommand) Usage() string { return "/channel <channel_name> <message>" }

func (c *ChannelCommand) BaseErrorMessage() string { return "Error sending message to channel" }

func (c *ChannelCommand) Execute(h *Hub, cmd Command) {

	if !h.RequireAuth(cmd, c.BaseErrorMessage()) {
		return
	}

	args, ok := h.GetArgs(cmd, 2, c.Usage(), c.BaseErrorMessage())
	if !ok {
		return
	}

	name := args[0]

	existingChannel := getChannelByName(h.channels, name)
	if existingChannel == nil {
		h.sendSystemToClient(
			cmd.From,
			"Can not send message. Channel with name '"+name+"' does not exist",
		)

		return
	}

	if existingChannel.Members[cmd.From.ID] == nil {
		h.sendSystemToClient(
			cmd.From,
			"Can not send message. You are not a member of channel '"+name+"'",
		)

		return
	}

	msg := Message{
		Text:    args[1],
		From:    cmd.From,
		Channel: existingChannel,
		Type:    MessageChannel,
	}

	h.handleSend(msg)
}
