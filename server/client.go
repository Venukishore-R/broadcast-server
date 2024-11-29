package server

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

type Client struct {
	Socket *websocket.Conn

	Receive chan []byte

	Username string
}

func Read(conn *websocket.Conn) {
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("unable to read message: %v\n", err)
			os.Exit(0)
			return
		}
		fmt.Println(string(msg))
	}
}

func Write(conn *websocket.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for {

		if scanner.Scan() {
			msg := scanner.Text()
			err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Printf("unable to send message: %v\n", err)
				return
			}
		} else if err := scanner.Err(); err != nil {
			log.Printf("Error reading input: %v\n", err)
			os.Exit(0)
			return
		}
	}

}

func StartClient(port string, userName string) {
	serverURL := fmt.Sprintf("ws://localhost:%s/ws/%s", port, userName)

	conn, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
	if err != nil {
		fmt.Printf("Error connecting to server: %v\n", err)
		return
	}

	defer conn.Close()

	go Read(conn)

	Write(conn)
}
