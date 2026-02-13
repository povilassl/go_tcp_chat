package hub

import (
	"strings"
)

type JoinCommand struct{}

func (c *JoinCommand) Name() string { return "join" }

func (c *JoinCommand) Usage() string { return "/join <channel_name>" }

func (c *JoinCommand) Execute(h *Hub, cmd Command) {
	if cmd.Args == "" {
		h.sendSystemToClient(
			cmd.From,
			"Incorrect number of arguments. Usage: /join <channel_name>",
		)

		return
	}

	name := strings.TrimSpace(cmd.Args)

	existingChannel := getChannelByName(h.channels, name)
	if existingChannel == nil {
		h.sendSystemToClient(
			cmd.From,
			"Channel with name '"+name+"' does not exist",
		)

		return
	}

	existingChannel.Members[cmd.From.ID] = cmd.From

	h.sendSystemToChannel(existingChannel, cmd.From.Name+" has joined the channel #"+name)
}
