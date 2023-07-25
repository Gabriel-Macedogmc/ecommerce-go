package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Gabriel-Macedogmc/products-backend/cmd/orderapi/router"
	"github.com/Gabriel-Macedogmc/products-backend/infra"
	"github.com/Gabriel-Macedogmc/products-backend/internal/domain/repository"
	"github.com/Gabriel-Macedogmc/products-backend/internal/domain/services"
)

const (
	rabbitMQURL  = "amqp://guest:guest@rabbitmq:5672/"
	queueName    = "my_queue"
	maxRetry     = 10
	retryTimeout = 5 * time.Second
)

func main() {
	db, err := infra.NewDBConnection()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	defer db.Statement.ReflectValue.Close()

	if err := waitForRabbitMQ(); err != nil {
		log.Fatalf("Failed to wait for RabbitMQ: %s", err)
	}

	mq, errMQ := infra.NewRabbitMQ("amqp://guest:guest@rabbitmq:5672/")
	if errMQ != nil {
		log.Fatalf("Error connecting to RabbitMQ: %v", errMQ)
	}

	productRepo := repository.NewProductRepository(db)
	productService := services.NewProductService(*productRepo, mq)

	router := router.NewRouter(*productService)

	router.Run(":3000")
	defer mq.Close()
}

func waitForRabbitMQ() error {
	log.Println("Waiting for RabbitMQ to be available...")

	for i := 0; i < maxRetry; i++ {
		// Tentar se conectar ao RabbitMQ
		conn, err := infra.NewRabbitMQ(rabbitMQURL)
		if err == nil {
			conn.Close()
			log.Println("RabbitMQ is available!")
			return nil
		}

		log.Printf("Failed to connect to RabbitMQ: %s", err)

		// Esperar um tempo antes de tentar novamente
		time.Sleep(retryTimeout)
	}

	return fmt.Errorf("RabbitMQ did not become available within the specified time")
}
