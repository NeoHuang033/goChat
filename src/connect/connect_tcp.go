package connect

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func (c *Connect) InitTcpServer() error {
	listenAddr := "0.0.0.0:8080"
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Printf("Error creating listener: %s\n", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Printf("Listening on %s\n", listenAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %s\n", err)
			continue
		}
		go handleConnection(conn)
	}

	return nil
}

func handleConnection(conn net.Conn) {
	//defer conn.Close()

	messages := "Hello, World!\n"
	_, err := conn.Write([]byte(messages))
	if err != nil {
		fmt.Printf("Error sending message: %s\n", err)
		return
	}
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading from connection: %s\n", err)
			return
		}
		fmt.Printf("Received message: %s", message)
	}
}
