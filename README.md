# TCP Chat

Welcome to the TCP Chat repository! This project showcases a robust and scalable TCP-based chat server built using Go, offering a unique approach to real-time communication over the traditional WebSockets protocol.

## Available features
- **Room-based Chat**: Organize conversations into specific rooms, enhancing topic-focused discussions.

- **Custom Nicknames**: Allow users to personalize their experience with custom nicknames.

- **Dynamic Room Management**: List and join available chat rooms dynamically.

- **Graceful Handling**: Proper client disconnection and error handling for a smooth user experience.

## Upcoming features
- User Authentication
- Private messaging
- Lobby
- Message history
- File transfer
- Search functionality
- End-To-End Encryption
- Rate limiting
- User moderation tools

## Benefits

- **Simplicity and Efficiency**: TCP sockets provide a simpler and more efficient way to manage connections compared to WebSockets, reducing overhead and improving performance.

- **Lightweight**: TCP chat applications are generally lighter in terms of resource usage, making them ideal for scenarios with limited computational resources.

- **Scalability**: With Go's powerful concurrency model, the server can handle numerous connections efficiently, making it highly scalable.

- **Security**: TCP can be combined with TLS to provide a secure communication channel, ensuring data integrity and confidentiality.

- **Flexibility**: Unlike WebSockets, which are tied to HTTP, TCP chat applications can operate independently, offering greater flexibility in deployment and integration.

## Why Choose TCP over WebSockets?

While WebSockets are excellent for integrating real-time communication into web applications, TCP offers several distinct advantages:

- **Direct Connection Management**: TCP provides direct control over the connection, allowing for more customized handling of the communication process.

- **Reduced Latency**: TCP connections can reduce latency due to the absence of HTTP overhead, leading to faster message delivery.

- **Protocol Agnostic**: TCP isn't bound to the constraints of the HTTP protocol, allowing for more versatile application design.

## Getting Started

### Prerequisites

- Go (version 1.18 or later) 

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/zde37/TCP-Chat.git
   cd TCP-Chat
   ```

2. Build the application:
   ```sh
   go build -o chat .
   ```

### Usage

1. Run the server:

   ```sh
   ./chat
   ```

2. Connect to the server using `telnet` or `nc`:
   ```sh
   telnet localhost 7000
   ```
   ```sh
   nc localhost 7000
   ```

### Commands

- `/nick [name]` - Set your nickname.
- `/join [room]` - Join a chat room(room name must begin with '#').
- `/rooms` - List available chat rooms.
- `/msg [message]` - Send a message to the current chat room.
- `/quit` - Disconnect from the server.

### Example

```sh
/nick Alice

/join #general

/msg Hello, everyone!

/rooms

/quit
```

## Contributing

1. Fork the repository.
2. Create a new branch: git checkout -b my-feature
3. Make your changes and commit them: git commit -m 'Add some feature'
4. Push to the branch: git push origin my-feature
5. Open a pull request.
6. Don't forget to star the repository and tell people about it.