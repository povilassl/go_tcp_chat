package hub

import (
	"fmt"
	"slices"
)

type ChannelCommand struct{}

func (c *ChannelCommand) Name() string { return "channel" }

func (c *ChannelCommand) Usage() string { return "/channel <channel_name> <message>" }

func (c *ChannelCommand) BaseErrorMessage() string { return "Error sending message to channel" }

func (c *ChannelCommand) Execute(h *Hub, cmd Command) {
	if !h.requireAuth(cmd, c.BaseErrorMessage()) {
		return
	}

	args, ok := h.getArgs(cmd, 2, c.Usage(), c.BaseErrorMessage())
	if !ok {
		return
	}

	channelName := args[0]
	messageText := args[1]

	channel, err := h.channelService.GetByName(channelName)
	if err != nil || channel == nil {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: Channel '%s' does not exist", c.BaseErrorMessage(), channelName),
		)
		return
	}

	members, err := h.channelService.GetMembers(channel.ID)
	if err != nil {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: %s", c.BaseErrorMessage(), err.Error()),
		)
		return
	}

	isMember := slices.Contains(*members, cmd.From.UserID)

	if !isMember {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: You are not a member of channel '%s'", c.BaseErrorMessage(), channelName),
		)
		return
	}

	if _, err := h.messageService.Create(cmd.From.UserID, nil, &channel.ChannelName, messageText); err != nil {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: %s", c.BaseErrorMessage(), err.Error()),
		)
		return
	}

	msg := Message{
		Text:        messageText,
		From:        cmd.From,
		ChannelName: channel.ChannelName,
		Type:        MessageChannel,
	}

	h.sendToUserIDs(*members, msg, &cmd.From.UserID)
}
