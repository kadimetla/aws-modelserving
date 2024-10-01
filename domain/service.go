package domain

import "log"

// MessageProcessor is a domain service that processes SQS messages.
type MessageProcessor interface {
    ProcessMessage(message SQSMessage) error
}

// SimpleMessageProcessor is an example of a basic implementation of MessageProcessor.
type SimpleMessageProcessor struct{}

func (s SimpleMessageProcessor) ProcessMessage(message SQSMessage) error {
    log.Printf("Processing message ID: %s with body: %s", message.ID, message.Body)
    // Add core business logic here
    return nil
}
