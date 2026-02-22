package hub

import (
	"fmt"
)

type LoginCommand struct{}

func (c *LoginCommand) Name() string { return "login" }

func (c *LoginCommand) Usage() string { return "/login <username> <password>" }

func (c *LoginCommand) BaseErrorMessage() string { return "Error logging in" }

func (c *LoginCommand) Execute(h *Hub, cmd Command) {
	args, ok := h.getArgs(cmd, 2, c.Usage(), c.BaseErrorMessage())
	if !ok {
		return
	}

	name := args[0]
	pass := args[1]

	user, err := h.authService.Login(name, pass)
	if err != nil {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: %s", c.BaseErrorMessage(), err.Error()),
		)
		return
	}

	if err := h.bindClientUser(cmd.From, user); err != nil {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: %s", c.BaseErrorMessage(), err.Error()),
		)
		return
	}

	h.sendSystemToClient(
		cmd.From,
		fmt.Sprintf("Successfully logged in as %s.", user.Nickname),
	)
}
