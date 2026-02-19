package hub

import (
	"fmt"
	"strings"
)

type LoginCommand struct{}

func (c *LoginCommand) Name() string { return "login" }

func (c *LoginCommand) Usage() string { return "/login <username> <password>" }

func (c *LoginCommand) BaseErrorMessage() string { return "Error logging in" }

func (c *LoginCommand) Execute(h *Hub, cmd Command) {
	args, ok := h.GetArgs(cmd, 2, c.Usage(), c.BaseErrorMessage())
	if !ok {
		return
	}

	name := strings.TrimSpace(args[0])
	pass := strings.TrimSpace(args[1])

	user, err := h.authService.Login(name, pass)
	if err != nil {
		h.sendSystemToClient(
			cmd.From,
			err.Error(),
		)

		return
	}

	cmd.From.Login(user)

	h.sendSystemToClient(
		cmd.From,
		fmt.Sprintf("Welcome to the server, %s!", cmd.From.User.Nickname),
	)
}
