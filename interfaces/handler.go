package interfaces

import (
    "context"
    "github.com/yourusername/sqs-reader/application"
    "log"
)

// SQSHandler starts the SQS message processing.
type SQSHandler struct {
    processor *application.SQSProcessor
}

func NewSQSHandler(processor *application.SQSProcessor) *SQSHandler {
    return &SQSHandler{
        processor: processor,
    }
}

func (h *SQSHandler) Start(ctx context.Context) {
    log.Println("Starting SQS Handler...")
    h.processor.StartProcessing(ctx)
}
