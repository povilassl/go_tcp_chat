package hub

import "fmt"

type RegisterCommand struct{}

func (c *RegisterCommand) Name() string { return "register" }

func (c *RegisterCommand) Usage() string { return "/register <username> <password> [nickname]" }

func (c *RegisterCommand) BaseErrorMessage() string { return "Error registering user" }

func (c *RegisterCommand) Execute(h *Hub, cmd Command) {
	args, ok := h.getArgsRange(cmd, 2, 3, c.Usage(), c.BaseErrorMessage())
	if !ok {
		return
	}

	name := args[0]
	pass := args[1]
	var nick *string

	if len(args) == 3 {
		nick = &args[2]
	}

	err := h.authService.Register(name, pass, nick)
	if err != nil {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: %s", c.BaseErrorMessage(), err.Error()),
		)

		return
	}

	h.sendSystemToClient(
		cmd.From,
		fmt.Sprintf("Successfully registered user %s. You can now log in using '/login %s <password>'.", name, name),
	)
}
