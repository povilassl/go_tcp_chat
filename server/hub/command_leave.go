package hub

import (
	"strings"
)

type LeaveCommand struct{}

func (c *LeaveCommand) Name() string { return "leave" }

func (c *LeaveCommand) Execute(h *Hub, cmd Command) {
	if cmd.Args == "" {
		h.sendSystem(
			cmd.From,
			"Incorrect number of arguments. Usage: /leave <channel_name>",
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

	delete(existingChannel.Members, cmd.From.ID)

	// TODO: send message to channel that user has left
	// TODO: send message to user that they have left the channel
}
