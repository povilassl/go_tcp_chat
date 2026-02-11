package hub

import (
	"strings"
)

type LeaveCommand struct{}

func (c *LeaveCommand) Name() string { return "leave" }

func (c *LeaveCommand) Usage() string { return "/leave <channel_name>" }

func (c *LeaveCommand) Execute(h *Hub, cmd Command) {
	if cmd.Args == "" {
		h.sendSystemToClient(
			cmd.From,
			"Incorrect number of arguments. Usage: /leave <channel_name>",
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

	delete(existingChannel.Members, cmd.From.ID)

	h.sendSystemToChannel(existingChannel, cmd.From.Name+" has left the channel #"+name)
}
