package infra

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Gabriel-Macedogmc/orders-backend/internal/domain/model"
	"github.com/Gabriel-Macedogmc/orders-backend/internal/domain/services"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	service *services.OrderService
}

func NewRabbitMQ(connStr string, service *services.OrderService) (*RabbitMQ, error) {
	conn, err := amqp.Dial(connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %s", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open RabbitMQ channel: %s", err)
	}

	consumer := &RabbitMQ{
		conn:    conn,
		channel: channel,
		service: service,
	}

	return consumer, nil
}
func (r *RabbitMQ) Consume() error {
	log.Printf("Consuming messages")
	msgs, err := r.channel.Consume(
		"orders",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to register RabbitMQ consumer: %s", err)
	}

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			log.Printf("Received messages")
			order := model.Order{}
			if err := json.Unmarshal(msg.Body, &order); err != nil {
				log.Printf("Error handling message: %s", err)
				msg.Reject(false)
				continue
			}

			if err := r.service.CreateOrder(order); err != nil {
				log.Printf("Failed to create order: %s", err)
				msg.Reject(false)
				continue
			}

			msg.Ack(false)
			log.Printf("Order created")
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit, press CTRL+C")
	<-forever

	return nil
}

func (r *RabbitMQ) Close() error {
	log.Printf("Closing RabbitMQ connection")
	if r.conn != nil {
		err := r.conn.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
