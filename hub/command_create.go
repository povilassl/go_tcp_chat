package hub

import (
	"fmt"
	"strings"
)

type CreateCommand struct{}

func (c *CreateCommand) Name() string { return "create" }

func (c *CreateCommand) Usage() string { return "/create <channel_name>" }

func (c *CreateCommand) BaseErrorMessage() string { return "Error creating channel" }

func (c *CreateCommand) Execute(h *Hub, cmd Command) {
	if !h.RequireAuth(cmd, c.BaseErrorMessage()) {
		return
	}

	args, ok := h.GetArgs(cmd, 1, c.Usage(), c.BaseErrorMessage())
	if !ok {
		return
	}

	name := strings.TrimSpace(args[0])

	err := h.channelService.Create(name, cmd.From.User)
	if err != nil {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: %s", c.BaseErrorMessage(), err.Error()),
		)
		return
	}

	h.sendSystemToClient(
		cmd.From,
		fmt.Sprintf("Channel '%s' created successfully", name),
	)
}
