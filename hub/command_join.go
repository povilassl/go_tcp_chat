package hub

import (
	"strings"
)

type JoinCommand struct{}

func (c *JoinCommand) Name() string { return "join" }

func (c *JoinCommand) Usage() string { return "/join <channel_name>" }

func (c *JoinCommand) BaseErrorMessage() string { return "Error joining channel" }

func (c *JoinCommand) Execute(h *Hub, cmd Command) {
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

	existingChannel.Members[cmd.From.ID] = cmd.From

	h.sendSystemToChannel(existingChannel, cmd.From.Name+" has joined the channel #"+name)
}
