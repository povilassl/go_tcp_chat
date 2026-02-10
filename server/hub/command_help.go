package hub

type HelpCommand struct{}

func (c *HelpCommand) Name() string { return "help" }

func (c *HelpCommand) Execute(h *Hub, cmd Command) {
	h.sendSystem(
		cmd.From,
		"Available commands: TODO",
	)
}
