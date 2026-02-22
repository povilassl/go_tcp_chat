package hub

import (
	"fmt"
	"strings"
)

type DeleteCommand struct{}

func (c *DeleteCommand) Name() string { return "delete" }

func (c *DeleteCommand) Usage() string { return "/delete <channel_name>" }

func (c *DeleteCommand) BaseErrorMessage() string { return "Error deleting channel" }

func (c *DeleteCommand) Execute(h *Hub, cmd Command) {
	if !h.RequireAuth(cmd, c.BaseErrorMessage()) {
		return
	}

	args, ok := h.GetArgs(cmd, 1, c.Usage(), c.BaseErrorMessage())
	if !ok {
		return
	}

	user, err := h.userService.GetByID(cmd.From.UserID)
	if err != nil {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: %s", c.BaseErrorMessage(), err.Error()),
		)
		return
	}

	name := strings.TrimSpace(args[0])

	err = h.channelService.Delete(name, user)
	if err != nil {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: %s", c.BaseErrorMessage(), err.Error()),
		)
		return
	}

	h.sendSystemToClient(
		cmd.From,
		fmt.Sprintf("Channel '%s' deleted successfully", name),
	)
}
