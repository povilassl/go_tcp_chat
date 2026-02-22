# Go TCP Chat

A lightweight chat server in Go with persistent message history, user authentication, and channel management. Inspired by IRC and Discord.

# Running the server

1. Create `.env` from the example:
   ```bash
   cp .env.example .env
   ```

2. Update `.env` with your database credentials

3. Start the server:
   ```bash
   docker compose up --build
   ```

### Done

- [x] User authentication and channels
- [x] TCP hub with graceful shutdown
- [x] Persistent storage with MySQL
- [x] Docker support

### Backlog

- [ ] Support multiple calls for same command
- [ ] File transfer
- [ ] Implement client
  - [ ] Ping / pong
  - [ ] Set color
- [ ] Set keys for system messages
