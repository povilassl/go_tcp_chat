package hub

import (
	"fmt"
	"strings"
)

type NameCommand struct{}

func (c *NameCommand) Name() string { return "name" }

func (c *NameCommand) Usage() string { return "/name <new_name>" }

func (c *NameCommand) BaseErrorMessage() string { return "Error changing name" }

func (c *NameCommand) Execute(h *Hub, cmd Command) {
	if !h.RequireAuth(cmd, c.BaseErrorMessage()) {
		return
	}

	args, ok := h.GetArgs(cmd, 1, c.Usage(), c.BaseErrorMessage())
	if !ok {
		return
	}

	originalName := cmd.From.User.Nickname
	newName := strings.TrimSpace(args[0])

	err := h.userService.Rename(cmd.From.User, &newName)
	if err != nil {
		h.sendSystemToClient(cmd.From, err.Error())
		return
	}

	h.sendSystemGlobalBroadcast(fmt.Sprintf("%s is now known as %s", originalName, newName))
}
