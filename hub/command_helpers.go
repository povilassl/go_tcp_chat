package hub

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func (h *Hub) requireAuth(cmd Command, baseErrorMessage string) bool {

	if cmd.From.UserID == uuid.Nil {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: You must be logged in to perform this action", baseErrorMessage),
		)
		return false
	}

	return true
}

func (h *Hub) getArgs(
	cmd Command,
	paramCount int,
	usage string,
	baseErrorMessage string) ([]string, bool) {

	return h.getArgsRange(cmd, paramCount, paramCount, usage, baseErrorMessage)
}

func (h *Hub) getArgsRange(
	cmd Command,
	minParamCount int,
	maxParamCount int,
	usage string,
	baseErrorMessage string) ([]string, bool) {

	if minParamCount < 0 || maxParamCount < minParamCount {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: Invalid argument configuration", baseErrorMessage),
		)

		return []string{}, false
	}

	trimmedArgs := strings.TrimSpace(cmd.Args)
	if trimmedArgs == "" {
		if minParamCount == 0 {
			return []string{}, true
		}

		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: Incorrect number of arguments. Usage: %s", baseErrorMessage, usage),
		)

		return []string{}, false
	}

	args := strings.SplitN(trimmedArgs, " ", maxParamCount)
	for i := range args {
		args[i] = strings.TrimSpace(args[i])
	}

	if len(args) < minParamCount || len(args) > maxParamCount {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: Incorrect number of arguments. Usage: %s", baseErrorMessage, usage),
		)

		return []string{}, false
	}

	return args, true
}
