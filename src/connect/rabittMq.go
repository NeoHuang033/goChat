package connect

import (
	"github.com/streadway/amqp"
)

type RabbitMQClient struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   amqp.Queue
}

func NewRabbitMQClient(amqpURL, queueName string) (*RabbitMQClient, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQClient{
		Conn:    conn,
		Channel: ch,
		Queue:   q,
	}, nil
}

func (client *RabbitMQClient) PublishMessage(body string) error {
	err := client.Channel.Publish(
		"",
		client.Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (client *RabbitMQClient) Close() {
	if client.Channel != nil {
		client.Channel.Close()
	}
	if client.Conn != nil {
		client.Conn.Close()
	}
}
