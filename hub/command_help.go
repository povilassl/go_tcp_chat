package hub

type HelpCommand struct{}

func (c *HelpCommand) Name() string { return "help" }

func (c *HelpCommand) Usage() string { return "/help" }

func (c *HelpCommand) BaseErrorMessage() string { return "" }

func (c *HelpCommand) Execute(h *Hub, cmd Command) {

	message := "Available commands:\r\n"

	for _, cmd := range h.commands {
		if usageCmd, ok := cmd.(interface{ Usage() string }); ok {
			message += "- " + usageCmd.Usage() + "\r\n"
		}
	}

	h.sendSystemToClient(
		cmd.From,
		message,
	)
}
