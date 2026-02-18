package hub

import (
	"fmt"
	"strings"
)

func (h *Hub) RequireAuth(cmd Command, baseErrorMessage string) bool {

	if cmd.From.User == nil {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: You must be logged in to perform this action", baseErrorMessage),
		)
		return false
	}

	return true
}

func (h *Hub) GetArgs(
	cmd Command,
	paramCount int,
	usage string,
	baseErrorMessage string) ([]string, bool) {

	args := strings.SplitN(cmd.Args, " ", paramCount)

	//TODO validate if works when 1
	if len(args) != paramCount {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: Incorrect number of arguments. Usage: %s", baseErrorMessage, usage),
		)

		return []string{}, false
	}

	return args, true
}
