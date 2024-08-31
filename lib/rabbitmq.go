package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/rabbitmq/amqp091-go"
)

var (
	connection *amqp091.Connection
	channel    *amqp091.Channel
)

func InitializeRabbitMQ() error {
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	if rabbitMQURL == "" {
		return fmt.Errorf("RABBITMQ_URL environment variable not set")
	}

	var err error
	connection, err = amqp091.Dial(rabbitMQURL)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	channel, err = connection.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %w", err)
	}

	return nil
}

func GetChannel() *amqp091.Channel {
	if channel == nil {
		log.Println("Channel is not initialized")
		return nil
	}
	return channel
}

func CloseRabbitMQ() {
	if channel != nil {
		if err := channel.Close(); err != nil {
			log.Printf("Failed to close channel: %v", err)
		}
	}
	if connection != nil {
		if err := connection.Close(); err != nil {
			log.Printf("Failed to close connection: %v", err)
		}
	}
}

func SendVerificationEmail(email, token string) error {
	channel := GetChannel()
	if channel == nil {
		return fmt.Errorf("channel is not initialized")
	}
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "example"
	}
	message := map[string]interface{}{
		"email": map[string]string{
			"subject": fmt.Sprintf("Please verify your email - %s", appName),
			"content": fmt.Sprintf("Register berhasil, segera aktifasi akun anda dengan memasukan token <b>%s</b>", token),
			"from":    "admin@gmail.com",
			"to":      email,
		},
	}

	payload, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	err = channel.Publish(
		"notification",
		"email",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        payload,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}
