package hub

import (
	"strings"
)

type JoinCommand struct{}

func (c *JoinCommand) Name() string { return "join" }

func (c *JoinCommand) Execute(h *Hub, cmd Command) {
	if cmd.Args == "" {
		h.sendSystem(
			cmd.From,
			"Incorrect number of arguments. Usage: /join <channel_name>",
		)

		return
	}

	name := strings.TrimSpace(cmd.Args)

	existingChannel := getChannelByName(h.channels, name)
	if existingChannel == nil {
		h.sendSystem(
			cmd.From,
			"Channel with name '"+name+"' does not exist",
		)

		return
	}

	existingChannel.Members[cmd.From.ID] = cmd.From

	// TODO: send message to channel that user has joined
	// TODO: send message to user that they have joined the channel
}
