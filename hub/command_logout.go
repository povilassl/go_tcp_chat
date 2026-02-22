package hub

import (
	"fmt"

	"github.com/google/uuid"
)

type LogoutCommand struct{}

func (c *LogoutCommand) Name() string { return "logout" }

func (c *LogoutCommand) Usage() string { return "/logout" }

func (c *LogoutCommand) BaseErrorMessage() string { return "Error logging out" }

func (c *LogoutCommand) Execute(h *Hub, cmd Command) {
	if cmd.From.UserID == uuid.Nil {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: You are not logged in", c.BaseErrorMessage()),
		)
		return
	}

	displayName := cmd.From.DisplayName

	if err := h.unbindClientUser(cmd.From); err != nil {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: %s", c.BaseErrorMessage(), err.Error()),
		)
		return
	}

	h.sendSystemToClient(
		cmd.From,
		fmt.Sprintf("Successfully logged out from user %s.", displayName),
	)
}
