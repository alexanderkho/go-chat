# System Architecture

## User Stories

- As a user i want to connect to a chat room and chat with other users
- As a user i want to see messages from other users in the chat room
- As a user i want to send messages to the chat room
- As a user, i want my messages to persist for as long as someone is active in the chat room

## High level design

- Use websockets to connect to the chat room
- Use a database to store messages
- Use a server to handle the websocket connections
- when a user sends a message, the server will store the message in the database and broadcast the message to all connected users
- we get ephemeral behaviour by default as users will only see messages that were sent while they were connected to the chat room
- we don't even need to store the messages in the database, we can just store them in memory and broadcast them to all connected users

## APIs

- POST /messages
  - Request: { message: string }
  - Response: { message: string, timestamp: string }

## Database Schema

- users
  - id: int
  - username: string
  - password: string
  - created_at: datetime
  - updated_at: datetime
- messages
  - id: int
  - user_id: int
  - message: string
  - created_at: datetime
  - updated_at: datetime
