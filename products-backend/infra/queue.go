package infra

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn  *amqp.Connection
	ch    *amqp.Channel
	queue amqp.Queue
}

func NewRabbitMQ(url string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, errCH := conn.Channel()
	if errCH != nil {
		log.Fatalf("Failed to Open a channel: %s", errCH)
	}

	queue, errQueue := ch.QueueDeclare(
		"orders",
		true,
		false,
		false,
		false,
		nil,
	)

	if errQueue != nil {
		log.Fatalf("Failed to declare a queue: %s", errQueue)
	}

	return &RabbitMQ{conn, ch, queue}, nil
}

func (r *RabbitMQ) Publish(exchange, key string, msg []byte) error {
	err := r.ch.Publish(
		exchange,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQ) Consume(queueName string, consumerName string) (<-chan amqp.Delivery, error) {
	msgs, err := r.ch.Consume(
		queueName,
		consumerName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

func (r *RabbitMQ) Close() error {
	log.Printf("Closing rabbitmq")
	if r.conn != nil {
		err := r.conn.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
