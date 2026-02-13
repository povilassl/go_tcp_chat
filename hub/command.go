package hub

type Command struct {
	Name string
	Args string
	From *Client
}

type CommandHandler interface {
	Name() string
	Usage() string
	Execute(h *Hub, cmd Command)
}

// switch {
// case strings.HasPrefix(cmd.Text, "/name"):
// 	args := strings.SplitN(cmd.Text, " ", 2)
// 	if len(args) != 2 {
// 		h.handleSend(Message{
// 			Text: "Incorrect number of arguments. Usage: /name <new_name>",
// 			To:   cmd.From,
// 			Type: MessageSystem,
// 		})

// 		return
// 	}

// 	originalName := cmd.From.Name
// 	newName := strings.TrimSpace(args[1])

// 	// Validate new name
// 	if len(newName) == 0 || len(newName) > 14 {
// 		h.handleSend(Message{
// 			Text: "Name must be between 1 and 14 characters long",
// 			To:   cmd.From,
// 			Type: MessageSystem,
// 		})
// 		return
// 	}

// 	if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(newName) {
// 		h.handleSend(Message{
// 			Text: "Name must contain only letters and numbers",
// 			To:   cmd.From,
// 			Type: MessageSystem,
// 		})
// 		return
// 	}

// 	cmd.From.Rename(newName)

// 	h.handleBroadcast(Message{
// 		Text: fmt.Sprintf("%s is now known as %s", originalName, newName),
// 		Type: MessageSystem,
// 	})

// case strings.HasPrefix(cmd.Text, "/msg"):
// 	args := strings.SplitN(cmd.Text, " ", 3)
// 	if len(args) != 3 {
// 		h.handleSend(Message{
// 			Text: "Incorrect number of arguments. Usage: /msg <name> <message>",
// 			To:   cmd.From,
// 			Type: MessageSystem,
// 		})

// 		return
// 	}

// 	var clientName = args[1]
// 	var messageText = args[2]

// 	existingClient := h.findClientByName(clientName)
// 	if existingClient == nil {
// 		h.handleSend(Message{
// 			Text: "Client '" + clientName + "' is not currently online",
// 			To:   cmd.From,
// 			Type: MessageSystem,
// 		})
// 	}

// 	h.handleSend(Message{
// 		Text: messageText,
// 		To:   existingClient,
// 		From: cmd.From,
// 		Type: MessageDirect,
// 	})

// case strings.HasPrefix(cmd.Text, "/join"):
// 	args := strings.SplitN(cmd.Text, " ", 2)

// 	if len(args) != 2 {
// 		h.handleSend(Message{
// 			Text: "Incorrect number of arguments. Usage: /join <channel_name>",
// 			To:   cmd.From,
// 			Type: MessageSystem,
// 		})

// 		return
// 	}

// case strings.HasPrefix(cmd.Text, "/quit"):
// 	args := strings.SplitN(cmd.Text, " ", 2)

// 	exitMessage := fmt.Sprintf("%s left the server", cmd.From.Name)

// 	if len(args) == 2 {
// 		exitMessage += fmt.Sprintf(". Goodbye message: %s", args[1])
// 	}

// 	h.handleBroadcast(Message{
// 		Text: exitMessage,
// 		Type: MessageSystem,
// 	})

// 	h.handleDisconnect(DisconnectEvent{
// 		Client: cmd.From,
// 		Reason: "requested",
// 	})

// case strings.HasPrefix(cmd.Text, "/help"):
// 	h.handleSend(Message{
// 		Text: "Available commands: TODO",
// 		To:   cmd.From,
// 		Type: MessageSystem,
// 	})

// default:
// 	h.handleSend(Message{
// 		Text: "Unknown command, send '/help' for a list of available commands",
// 		To:   cmd.From,
// 		Type: MessageSystem,
// 	})
// }
