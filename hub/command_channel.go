package hub

import "fmt"

type ChannelCommand struct{}

func (c *ChannelCommand) Name() string { return "channel" }

func (c *ChannelCommand) Usage() string { return "/channel <channel_name> <message>" }

func (c *ChannelCommand) BaseErrorMessage() string { return "Error sending message to channel" }

func (c *ChannelCommand) Execute(h *Hub, cmd Command) {
	if !h.RequireAuth(cmd, c.BaseErrorMessage()) {
		return
	}

	args, ok := h.GetArgs(cmd, 2, c.Usage(), c.BaseErrorMessage())
	if !ok {
		return
	}

	channelName := args[0]
	messageText := args[1]

	existingChannel := getChannelByName(h.channels, channelName)
	if existingChannel == nil {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: Channel '%s' does not exist", c.BaseErrorMessage(), channelName),
		)

		return
	}

	if existingChannel.Members[cmd.From.ID] == nil {
		h.sendSystemToClient(
			cmd.From,
			fmt.Sprintf("%s: You are not a member of channel '%s'", c.BaseErrorMessage(), channelName),
		)

		return
	}

	msg := Message{
		Text:    messageText,
		From:    cmd.From,
		Channel: existingChannel,
		Type:    MessageChannel,
	}

	h.handleSend(msg)
}
