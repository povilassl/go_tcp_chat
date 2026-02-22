package hub

import (
	"fmt"
	"strings"
)

type HelpCommand struct{}

func (c *HelpCommand) Name() string { return "help" }

func (c *HelpCommand) Usage() string { return "/help" }

func (c *HelpCommand) BaseErrorMessage() string { return "" }

func (c *HelpCommand) Execute(h *Hub, cmd Command) {

	var message strings.Builder
	message.WriteString("Available commands:\r\n")

	for _, cmd := range h.commands {
		if usageCmd, ok := cmd.(interface{ Usage() string }); ok {
			message.WriteString(fmt.Sprintf(" - %s\r\n", usageCmd.Usage()))
		}
	}

	h.sendSystemToClient(
		cmd.From,
		message.String(),
	)
}
