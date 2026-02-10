# TCP chat

### Done

- [x] Commands
  - [x] Change name
  - [x] Quit
  - [x] Send informational system messages on action
  - [x] Direct messages
- [x] Basic logging
- [x] Graceful shutdown

### In progress

- [ ] Add channels
  - [x] Create
    - [ ] Change index from int to string?
  - [x] Delete
  - [x] Join
  - [x] Leave
  - [ ] Send message
  - [ ] Get list
  - [ ] Remove general broadcast
    - [ ] Direct all actions thru commands
    - [ ] Force join a main channel

### Backlog

- [ ] Add excluded clients to broadcast?
- [ ] Help message
- [ ] Persistence
  - [ ] Load messages?
  - [ ] Register / login
  - [ ] Limit creation of channels
- [ ] Support multiple calls for same command?
- [ ] File transfer?
- [ ] Implement client
  - [ ] Set color
- [ ] Set keys for system messages

#### Open thoughts

So we have 3 different ways of sending a message:

- Direct
- Public
- Channel

I think we should abolish the ability to send public messages
That way we are gonna be left with only two options

Maybe we can simply format them using brackets?

- when receiving directly - (<client_name>) <message>
- when receiving from channel - [<channel_name>] <client_name>: <message>

this way we would be able to see messages formatted nicely in plain text when viewing telnet
but also format them when accessing on our own client
