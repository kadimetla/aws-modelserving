package sqs

import (
    "context"
    "log"

    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/sqs"
    "github.com/yourusername/sqs-reader/domain"
)

type SQSClient struct {
    client   *sqs.Client
    queueURL string
}

// NewSQSClient creates a new SQSClient.
func NewSQSClient(queueURL string) (*SQSClient, error) {
    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
        return nil, err
    }
    return &SQSClient{
        client:   sqs.NewFromConfig(cfg),
        queueURL: queueURL,
    }, nil
}

// ReceiveMessages fetches messages from SQS.
func (s *SQSClient) ReceiveMessages(ctx context.Context) ([]domain.SQSMessage, error) {
    output, err := s.client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
        QueueUrl:            &s.queueURL,
        MaxNumberOfMessages: 10,
        WaitTimeSeconds:     20,
    })
    if err != nil {
        return nil, err
    }

    messages := make([]domain.SQSMessage, 0, len(output.Messages))
    for _, msg := range output.Messages {
        messages = append(messages, domain.SQSMessage{
            ID:     *msg.MessageId,
            Body:   *msg.Body,
            SentAt: time.Now(), // Assuming current time as message sent time
        })
    }
    return messages, nil
}

// DeleteMessage deletes a message from the SQS queue after processing.
func (s *SQSClient) DeleteMessage(ctx context.Context, receiptHandle *string) error {
    _, err := s.client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
        QueueUrl:      &s.queueURL,
        ReceiptHandle: receiptHandle,
    })
    return err
}
