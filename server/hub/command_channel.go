package hub

import (
	"strings"
)

type ChannelCommand struct{}

func (c *ChannelCommand) Name() string { return "channel" }

func (c *ChannelCommand) Execute(h *Hub, cmd Command) {
	args := strings.SplitN(cmd.Args, " ", 2)
	if len(args) != 2 {
		h.sendSystem(
			cmd.From,
			"Incorrect number of arguments. Usage: /channel <channel_name> <message>",
		)

		return
	}

	name := args[0]

	existingChannel := getChannelByName(h.channels, name)
	if existingChannel == nil {
		h.sendSystem(
			cmd.From,
			"Can not send message. Channel with name '"+name+"' does not exist",
		)

		return
	}

	if existingChannel.Members[cmd.From.ID] == nil {
		h.sendSystem(
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
