# Go TCP Chat

A lightweight TCP chat server in Go with persistent message history, user authentication, and channel management. Inspired by IRC and Discord.

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
