package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Define Kafka topic and partition
	topic := "my-topic"
	partition := 0

	// Function to create a new Kafka connection
	createKafkaConnection := func() *kafka.Conn {
		conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:29092", topic, partition)
		if err != nil {
			log.Fatalf("failed to dial leader: %v", err)
		}
		return conn
	}

	// Create initial Kafka connection
	conn := createKafkaConnection()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatalf("failed to close writer: %v", err)
		}
	}()

	// Create a new scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter messages to send to Kafka (type 'exit' to quit):")

	for {
		// Read user input
		if !scanner.Scan() {
			break
		}
		text := scanner.Text()

		// Exit the loop if the user types "exit"
		if text == "exit" {
			break
		}

		// Reset the write deadline before each write
		conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

		// Write the user input as a message to the Kafka topic
		_, err := conn.WriteMessages(
			kafka.Message{Key: []byte("user-message"), Value: []byte(text)},
		)
		if err != nil {
			log.Printf("failed to write message: %v", err)
			log.Println("reconnecting to Kafka...")
			conn = createKafkaConnection()
		} else {
			fmt.Println("Message sent:", text)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading input: %v", err)
	}
}
