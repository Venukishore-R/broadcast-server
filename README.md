# Broadcast Server

## Overview

The Broadcast Server is a simple command-line application that allows clients to connect and send messages that will be broadcasted to all connected clients. This project demonstrates real-time communication using WebSockets and helps in understanding the underlying principles of applications like chat systems and live scoreboards.

## Features

- Start a broadcast server that listens for client connections.
- Connect multiple clients to the server.
- Broadcast messages from any connected client to all other clients.
- Handle client connections and disconnections gracefully.
- Simple command-line interface for server and client operations.

## Requirements

- Go 1.16^ (or any programming language of your choice)
- `websockets` library (if using Python)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Venukishore-R/broadcast-server.git
   cd broadcast-server
   ```
2. Install the required libraries (if using Go):
   ```bash
   go mod tidy
   ```
3. Run the application:
   ```bash
   make run
   ```
   
## Usage

### Starting the Server

To start the broadcast server, use the following command:
  ```bash
  broadcast-server --start --port <PORT_NUMBER>
  ```
This command will start the server and listen for incoming client connections on a specified port.

### Connecting a Client

To connect a client to the server, use the following command:
  ```bash
  broadcast-server --connect --port <PORT_NUMBER>
  ```
Once connected, you can send messages to the server. The server will broadcast these messages to all connected clients.

### Sending Messages

After connecting as a client, type your message and press Enter to send it. All connected clients will receive the message.

### Disconnecting

To disconnect from the server, simply close the client application or use a specific command if implemented.

## Implementation Details

### Server

1. Create a server that listens for incoming WebSocket connections.
2. Maintain a list of connected clients.
3. Broadcast messages received from any client to all other connected clients.
4. Handle client disconnections and remove them from the list.

### Client

1. Implement a client that can connect to the server via WebSocket.
2. Allow users to send messages to the server.


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- Inspired by real-time communication applications and WebSocket protocols.
- Special thanks to the open-source community for libraries and resources.

## Contributing

Feel free to submit issues, fork the repository, and create pull requests for enhancements or bug fixes.

## Conclusion

For more details about the project, check out the [Task Tracker Project on Roadmap.sh](https://roadmap.sh/projects/broadcast-server).
