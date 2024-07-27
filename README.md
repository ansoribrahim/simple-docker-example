# Simple Kafka Example

This project sets up a Kafka environment using Docker Compose. It includes services for Zookeeper, Kafka, and Kafka UI, along with example Go applications for producing and consuming messages.

## Docker Compose Setup

### `docker-compose.yml`

This file defines the following services:

- **zookeeper**: Provides Zookeeper service required by Kafka.
- **kafka**: Provides Kafka broker service.
- **kafka-ui**: Provides a web UI for managing and monitoring Kafka.

### Configuration

- **Zookeeper**: Exposes port 2181.
- **Kafka**: Exposes ports 9092 (internal) and 29092 (external for local access).
- **Kafka UI**: Exposes port 8080 for web access.

## Running the Setup

1. **Start the Services**

   ```bash
   docker-compose up
   ```
This command will start all services defined in the docker-compose.yml file.

2. **update go modules**

   ```bash
   go mod tidy
   ```
This command will update the go modules to the latest versions.

3. Access Kafka UI

Open your web browser and navigate to http://localhost:8080 to access Kafka UI.

## Go Applications

## Consumer
The consumer/main.go file contains a simple Kafka consumer that reads messages from the my-topic topic.

### To run the consumer:

Ensure Kafka and Zookeeper are running.
run the consumer application:

```bash
go run consumer/main.go
```

## Producer
The producer/main.go file contains a simple Kafka producer that allows you to send messages to the my-topic topic via standard input.

### To run the producer:

Ensure Kafka and Zookeeper are running.
run the producer application:

```bash
go run producer/main.go
```

Producer will send messages to the my-topic topic via standard input. 
You can run both the consumer and producer applications, and see the messages you input in the producer application will be displayed in the consumer application.
