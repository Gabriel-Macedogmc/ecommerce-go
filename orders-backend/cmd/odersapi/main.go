package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Gabriel-Macedogmc/orders-backend/infra"
	"github.com/Gabriel-Macedogmc/orders-backend/internal/domain/repository"
	"github.com/Gabriel-Macedogmc/orders-backend/internal/domain/services"
	"github.com/streadway/amqp"
)

const (
	rabbitMQURL  = "amqp://guest:guest@rabbitmq:5672/"
	maxRetry     = 10
	retryTimeout = 5 * time.Second
)

func main() {
	db, err := infra.NewDBConnection()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	cnn, _ := db.DB()
	defer cnn.Close()

	orderRepository := repository.NewOrderRepository(db)
	service := services.NewOrderService(*orderRepository)

	if err := waitForRabbitMQ(); err != nil {
		log.Fatalf("Failed to wait for RabbitMQ: %s", err)
	}

	mq, errMQ := infra.NewRabbitMQ(rabbitMQURL, service)
	if errMQ != nil {
		log.Fatalf("Error connecting to RabbitMQ: %v", errMQ)
	}

	fmt.Println("Running...")

	err = mq.Consume()
	if err != nil {
		log.Fatalf("Error consuming:  %v", err)
	}

	defer mq.Close()
}

func waitForRabbitMQ() error {
	log.Println("Waiting for RabbitMQ to be available...")

	for i := 0; i < maxRetry; i++ {
		conn, err := amqp.Dial(rabbitMQURL)
		if err == nil {
			conn.Close()
			log.Println("RabbitMQ is available!")
			return nil
		}

		log.Printf("Failed to connect to RabbitMQ: %s", err)

		time.Sleep(retryTimeout)
	}

	return fmt.Errorf("RabbitMQ did not become available within the specified time")
}
