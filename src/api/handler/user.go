package api

import (
	"encoding/json"
	//"bufio"
	//"encoding/json"
	"fmt"
	"goChat/src/connect"
	"goChat/src/domain"
	"log"

	//"io"
	"net"
	"strings"
	"sync"
)

type UserConn struct {
	UserName string
	Conn     net.Conn
	//Online   bool
}

type ConnManager struct {
	Conncetions map[string]*UserConn
	Mutex       sync.Mutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		Conncetions: make(map[string]*UserConn),
	}
}

func (manager *ConnManager) AddConnection(userName string, conn net.Conn) {
	manager.Mutex.Lock()
	defer manager.Mutex.Unlock()
	manager.Conncetions[userName] = &UserConn{
		UserName: userName,
		Conn:     conn,
	}
}

func (manager *ConnManager) RemoveConnection(userName string, conn net.Conn) {
	manager.Mutex.Lock()
	defer manager.Mutex.Unlock()
	if conn, ok := manager.Conncetions[userName]; ok {
		conn.Conn.Close()
		delete(manager.Conncetions, userName)
	}
}

func (manager *ConnManager) HandleConnection(conn net.Conn, msg domain.Message) {

	manager.AddConnection(msg.Username, conn)
	defer manager.RemoveConnection(msg.Username, conn)

	for {
		message := msg.Message
		message = strings.TrimSpace(message)

		if errMesg := manager.ProcessMessage(msg); errMesg != nil {
			fmt.Printf("Error processing message: %s\n", errMesg)
			continue
		}
	}
}

func (manager *ConnManager) ProcessMessage(msg domain.Message) error {
	/*parts := strings.SplitN(message, ":", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid message format")
	}*/
	receiver := msg.Receiver
	message := msg.Message

	//check conn for receiver
	manager.Mutex.Lock()
	receiverconn, ok := manager.Conncetions[receiver]
	manager.Mutex.Unlock()

	if !ok {
		return fmt.Errorf("receiver not found")
	}

	//send message
	_, err := receiverconn.Conn.Write([]byte(msg.Username + " says: " + message + "\n"))
	if err != nil {
		return fmt.Errorf("error sending message to receiver: %s", err)
	}

	return nil
}

func (manager *ConnManager) StartConsuming(rabbitMQClient *connect.RabbitMQClient) {
	msgs, err := rabbitMQClient.Channel.Consume(
		rabbitMQClient.Queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			var msg domain.Message
			if err := json.Unmarshal(d.Body, &msg); err != nil {
				log.Printf("Error unmarshaling message: %s", err)
				continue
			}

			// 处理消息
			// ...
		}
	}()
}
