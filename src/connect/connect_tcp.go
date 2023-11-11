package connect

import (
	"bufio"
	"encoding/json"
	"fmt"
	api "goChat/src/api/handler"
	"goChat/src/domain"
	"io"
	"net"
	"os"
)

var connManager = api.NewConnManager()

func (c *Connect) InitTcpServer() error {
	listenAddr := "0.0.0.0:8081"
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Printf("Error creating listener: %s\n", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Printf("Listening on %s\n", listenAddr)
	clients := make(map[string]net.Conn)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %s\n", err)
			continue
		}
		go handleConnection(conn, clients)
	}

	return nil
}

func handleConnection(conn net.Conn, clients map[string]net.Conn) {
	defer conn.Close()
	var msg domain.Message

	reader := bufio.NewReader(conn)

	message, err := reader.ReadString('\n')
	if err != nil {
		if err != io.EOF {
			fmt.Printf("Error reading from connection: %s\n", err)
		}
		return
	}
	/*token, _ := reader.ReadString('\n')
	token = strings.TrimSpace(token)*/

	if err := json.Unmarshal([]byte(message), &msg); err != nil {
		fmt.Printf("Error parsing JSON message: %s\n", err)
		return
	}

	if err, _ := VerifyToken(msg.Token); !err {
		fmt.Println("Invalid token. Closing connection.")
		return
	}
	go connManager.HandleConnection(conn, msg)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading from connection: %s\n", err)
			return
		}
		fmt.Printf("Received message: %s", message)
	}
}
