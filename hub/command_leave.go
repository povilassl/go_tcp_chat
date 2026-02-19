package hub

import (
	"fmt"
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
			fmt.Sprintf("%s: Channel '%s' does not exist", c.BaseErrorMessage(), name),
		)
		return
	}

	//TODO
	delete(existingChannel.Members, cmd.From.ID)

	h.sendSystemToChannel(
		existingChannel,
		fmt.Sprintf("%s has left the channel #%s", cmd.From.User.Nickname, name),
	)
}
