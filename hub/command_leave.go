package hub

import (
	"strings"
)

type LeaveCommand struct{}

func (c *LeaveCommand) Name() string { return "leave" }

func (c *LeaveCommand) Usage() string { return "/leave <channel_name>" }

func (c *LeaveCommand) BaseErrorMessage() string { return "Error leaving channel" }

func (c *LeaveCommand) Execute(h *Hub, cmd Command) {
	if !h.RequireAuth(cmd, c.BaseErrorMessage()) {
		return
	}

	args, ok := h.GetArgs(cmd, 1, c.Usage(), c.BaseErrorMessage())
	if !ok {
		return
	}

	name := strings.TrimSpace(args[0])

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
