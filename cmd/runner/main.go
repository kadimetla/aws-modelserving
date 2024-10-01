package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/yourusername/sqs-reader/application"
    "github.com/yourusername/sqs-reader/domain"
    "github.com/yourusername/sqs-reader/infrastructure/sqs"
    "github.com/yourusername/sqs-reader/interfaces"
)

func main() {
    queueURL := "https://sqs.us-east-1.amazonaws.com/123456789012/MyQueue"

    // Setup SQS Client
    sqsClient, err := sqs.NewSQSClient(queueURL)
    if err != nil {
        log.Fatalf("Failed to create SQS client: %v", err)
    }

    // Setup Domain Processor
    processor := domain.SimpleMessageProcessor{}

    // Setup SQS Processor
    sqsProcessor := application.NewSQSProcessor(sqsClient, processor)

    // Setup Interface Handler
    handler := interfaces.NewSQSHandler(sqsProcessor)

    // Handle graceful shutdown
    ctx, cancel := context.WithCancel(context.Background())
    c := make(chan os.Signal, 1)
    signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        <-c
        log.Println("Received shutdown signal")
        cancel()
    }()

    // Start the handler to process messages from SQS
    handler.Start(ctx)

    // Wait for graceful shutdown
    time.Sleep(1 * time.Second)
    log.Println("Shutdown complete")
}
