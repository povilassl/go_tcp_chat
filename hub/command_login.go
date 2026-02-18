package hub

import (
	"strings"
)

type LoginCommand struct{}

func (c *LoginCommand) Name() string { return "login" }

func (c *LoginCommand) Usage() string { return "/login <username> <password>" }

//TODO: continue adding base messages and require auth commands

func (c *LoginCommand) Execute(h *Hub, cmd Command) {
	args := strings.SplitN(cmd.Args, " ", 2)

	if len(args) != 2 {
		h.sendSystemToClient(
			cmd.From,
			"Incorrect number of arguments. Usage: /register <username> <password>",
		)

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
		"User '"+name+"' logged in successfully",
	)
}
